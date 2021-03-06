// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CreateUserProvidedServiceInstanceRequest create user provided service instance request
//
// swagger:model createUserProvidedServiceInstanceRequest
type CreateUserProvidedServiceInstanceRequest struct {

	// A hash that can be used to store credentials
	Credentials GenericObject `json:"credentials,omitempty"`

	// A name for the service instance
	Name string `json:"name,omitempty"`

	// The guid of the space in which the instance will be created
	SpaceGUID string `json:"space_guid,omitempty"`

	// The url for the syslog_drain to direct to
	SyslogDrainURL string `json:"syslog_drain_url,omitempty"`
}

// Validate validates this create user provided service instance request
func (m *CreateUserProvidedServiceInstanceRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCredentials(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateUserProvidedServiceInstanceRequest) validateCredentials(formats strfmt.Registry) error {

	if swag.IsZero(m.Credentials) { // not required
		return nil
	}

	if err := m.Credentials.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("credentials")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateUserProvidedServiceInstanceRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateUserProvidedServiceInstanceRequest) UnmarshalBinary(b []byte) error {
	var res CreateUserProvidedServiceInstanceRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
