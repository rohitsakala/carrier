// Code generated by go-swagger; DO NOT EDIT.

package security_group_staging_defaults

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ReturnSecurityGroupsUsedForStagingHandlerFunc turns a function with the right signature into a return security groups used for staging handler
type ReturnSecurityGroupsUsedForStagingHandlerFunc func(ReturnSecurityGroupsUsedForStagingParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReturnSecurityGroupsUsedForStagingHandlerFunc) Handle(params ReturnSecurityGroupsUsedForStagingParams) middleware.Responder {
	return fn(params)
}

// ReturnSecurityGroupsUsedForStagingHandler interface for that can handle valid return security groups used for staging params
type ReturnSecurityGroupsUsedForStagingHandler interface {
	Handle(ReturnSecurityGroupsUsedForStagingParams) middleware.Responder
}

// NewReturnSecurityGroupsUsedForStaging creates a new http.Handler for the return security groups used for staging operation
func NewReturnSecurityGroupsUsedForStaging(ctx *middleware.Context, handler ReturnSecurityGroupsUsedForStagingHandler) *ReturnSecurityGroupsUsedForStaging {
	return &ReturnSecurityGroupsUsedForStaging{Context: ctx, Handler: handler}
}

/*ReturnSecurityGroupsUsedForStaging swagger:route GET /config/staging_security_groups securityGroupStagingDefaults returnSecurityGroupsUsedForStaging

Return the Security Groups used for staging

curl --insecure -i %s/v2/config/staging_security_groups -X GET -H 'Authorization: %s'

*/
type ReturnSecurityGroupsUsedForStaging struct {
	Context *middleware.Context
	Handler ReturnSecurityGroupsUsedForStagingHandler
}

func (o *ReturnSecurityGroupsUsedForStaging) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReturnSecurityGroupsUsedForStagingParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
