package deployments

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/kyokomi/emoji"
	"github.com/pkg/errors"
	"github.com/suse/carrier/cli/helpers"
	"github.com/suse/carrier/cli/kubernetes"
	"github.com/suse/carrier/cli/paas/ui"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Registry struct {
	Debug   bool
	Timeout int
}

const (
	registryDeploymentID = "carrier-registry"
	registryVersion      = "0.1.0"
	registryChartFile    = "container-registry-0.1.0.tgz"
)

func (k *Registry) NeededOptions() kubernetes.InstallationOptions {
	return kubernetes.InstallationOptions{
		{
			Name:        "system_domain",
			Description: "The domain you are planning to use for Carrier. Should be pointing to the traefik public IP",
			Type:        kubernetes.StringType,
			Default:     "",
		},
	}
}

func (k *Registry) ID() string {
	return registryDeploymentID
}

func (k *Registry) Backup(c *kubernetes.Cluster, ui *ui.UI, d string) error {
	return nil
}

func (k *Registry) Restore(c *kubernetes.Cluster, ui *ui.UI, d string) error {
	return nil
}

func (k Registry) Describe() string {
	return emoji.Sprintf(":cloud:Registry version: %s\n:clipboard:Registry chart: %s", registryVersion, registryChartFile)
}

// Delete removes Registry from kubernetes cluster
func (k Registry) Delete(c *kubernetes.Cluster, ui *ui.UI) error {
	ui.Note().Msg("Removing Registry...")

	currentdir, err := os.Getwd()
	if err != nil {
		return errors.New("Failed uninstalling Registry: " + err.Error())
	}

	message := "Removing helm release " + registryDeploymentID
	out, err := helpers.WaitForCommandCompletion(ui, message,
		func() (string, error) {
			helmCmd := fmt.Sprintf("helm uninstall '%s' --namespace '%s'", registryDeploymentID, registryDeploymentID)
			return helpers.RunProc(helmCmd, currentdir, k.Debug)
		},
	)
	if err != nil {
		if strings.Contains(out, "release: not found") {
			ui.Exclamation().Msgf("%s helm release not found, skipping.\n", registryDeploymentID)
		} else {
			return errors.Wrapf(err, "Failed uninstalling helm release %s: %s", registryDeploymentID, out)
		}
	}

	message = "Deleting Registry namespace " + registryDeploymentID
	warning, err := helpers.WaitForCommandCompletion(ui, message,
		func() (string, error) {
			return c.DeleteNamespaceIfOwned(registryDeploymentID)
		},
	)
	if err != nil {
		return errors.Wrapf(err, "Failed deleting namespace %s", registryDeploymentID)
	}
	if warning != "" {
		ui.Exclamation().Msg(warning)
	}

	ui.Success().Msg("Registry removed")

	return nil
}

func (k Registry) apply(c *kubernetes.Cluster, ui *ui.UI, options kubernetes.InstallationOptions, upgrade bool) error {
	action := "install"
	if upgrade {
		action = "upgrade"
	}

	currentdir, err := os.Getwd()
	if err != nil {
		return err
	}

	if err = createQuarksMonitoredNamespace(c, registryDeploymentID); err != nil {
		return err
	}

	tarPath, err := helpers.ExtractFile(registryChartFile)
	if err != nil {
		return errors.New("Failed to extract embedded file: " + tarPath + " - " + err.Error())
	}
	defer os.Remove(tarPath)

	helmCmd := fmt.Sprintf("helm %s %s --create-namespace --namespace %s %s", action, registryDeploymentID, registryDeploymentID, tarPath)
	if out, err := helpers.RunProc(helmCmd, currentdir, k.Debug); err != nil {
		return errors.New("Failed installing Registry: " + out)
	}

	err = c.LabelNamespace(registryDeploymentID, kubernetes.CarrierDeploymentLabelKey, kubernetes.CarrierDeploymentLabelValue)
	if err != nil {
		return err
	}
	if err := c.WaitUntilPodBySelectorExist(ui, registryDeploymentID, "app.kubernetes.io/name=container-registry", 180); err != nil {
		return errors.Wrap(err, "failed waiting Registry deployment to come up")
	}
	if err := c.WaitForPodBySelectorRunning(ui, registryDeploymentID, "app.kubernetes.io/name=container-registry", 180); err != nil {
		return errors.Wrap(err, "failed waiting Registry deployment to come up")
	}

	ui.Success().Msg("Registry deployed")

	return nil
}

func (k Registry) GetVersion() string {
	return registryVersion
}

func (k Registry) Deploy(c *kubernetes.Cluster, ui *ui.UI, options kubernetes.InstallationOptions) error {

	_, err := c.Kubectl.CoreV1().Namespaces().Get(
		context.Background(),
		registryDeploymentID,
		metav1.GetOptions{},
	)
	if err == nil {
		return errors.New("Namespace " + registryDeploymentID + " present already")
	}

	ui.Note().Msg("Deploying Registry...")

	err = k.apply(c, ui, options, false)
	if err != nil {
		return err
	}

	return nil
}

func (k Registry) Upgrade(c *kubernetes.Cluster, ui *ui.UI, options kubernetes.InstallationOptions) error {
	_, err := c.Kubectl.CoreV1().Namespaces().Get(
		context.Background(),
		registryDeploymentID,
		metav1.GetOptions{},
	)
	if err != nil {
		return errors.New("Namespace " + registryDeploymentID + " not present")
	}

	ui.Note().Msg("Upgrading Registry...")

	return k.apply(c, ui, options, true)
}

func createQuarksMonitoredNamespace(c *kubernetes.Cluster, name string) error {
	_, err := c.Kubectl.CoreV1().Namespaces().Create(
		context.Background(),
		&corev1.Namespace{
			ObjectMeta: metav1.ObjectMeta{
				Name: name,
				Labels: map[string]string{
					"quarks.cloudfoundry.org/monitored": "quarks-secret",
				},
			},
		},
		metav1.CreateOptions{},
	)

	return err
}
