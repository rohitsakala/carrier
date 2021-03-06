// Code generated by go-swagger; DO NOT EDIT.

package service_plans

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// UpdateServicePlanDeprecatedHandlerFunc turns a function with the right signature into a update service plan deprecated handler
type UpdateServicePlanDeprecatedHandlerFunc func(UpdateServicePlanDeprecatedParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateServicePlanDeprecatedHandlerFunc) Handle(params UpdateServicePlanDeprecatedParams) middleware.Responder {
	return fn(params)
}

// UpdateServicePlanDeprecatedHandler interface for that can handle valid update service plan deprecated params
type UpdateServicePlanDeprecatedHandler interface {
	Handle(UpdateServicePlanDeprecatedParams) middleware.Responder
}

// NewUpdateServicePlanDeprecated creates a new http.Handler for the update service plan deprecated operation
func NewUpdateServicePlanDeprecated(ctx *middleware.Context, handler UpdateServicePlanDeprecatedHandler) *UpdateServicePlanDeprecated {
	return &UpdateServicePlanDeprecated{Context: ctx, Handler: handler}
}

/*UpdateServicePlanDeprecated swagger:route PUT /service_plans servicePlans updateServicePlanDeprecated

Updating a Service Plan (deprecated)

curl --insecure -i %s/v2/service_plans -X PUT -H 'Authorization: %s' -d '%s'

*/
type UpdateServicePlanDeprecated struct {
	Context *middleware.Context
	Handler UpdateServicePlanDeprecatedHandler
}

func (o *UpdateServicePlanDeprecated) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUpdateServicePlanDeprecatedParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
