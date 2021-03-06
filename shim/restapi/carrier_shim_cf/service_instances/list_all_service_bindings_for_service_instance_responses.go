// Code generated by go-swagger; DO NOT EDIT.

package service_instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/suse/carrier/shim/models"
)

// ListAllServiceBindingsForServiceInstanceOKCode is the HTTP code returned for type ListAllServiceBindingsForServiceInstanceOK
const ListAllServiceBindingsForServiceInstanceOKCode int = 200

/*ListAllServiceBindingsForServiceInstanceOK successful response

swagger:response listAllServiceBindingsForServiceInstanceOK
*/
type ListAllServiceBindingsForServiceInstanceOK struct {

	/*
	  In: Body
	*/
	Payload *models.ListAllServiceBindingsForServiceInstanceResponsePaged `json:"body,omitempty"`
}

// NewListAllServiceBindingsForServiceInstanceOK creates ListAllServiceBindingsForServiceInstanceOK with default headers values
func NewListAllServiceBindingsForServiceInstanceOK() *ListAllServiceBindingsForServiceInstanceOK {

	return &ListAllServiceBindingsForServiceInstanceOK{}
}

// WithPayload adds the payload to the list all service bindings for service instance o k response
func (o *ListAllServiceBindingsForServiceInstanceOK) WithPayload(payload *models.ListAllServiceBindingsForServiceInstanceResponsePaged) *ListAllServiceBindingsForServiceInstanceOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list all service bindings for service instance o k response
func (o *ListAllServiceBindingsForServiceInstanceOK) SetPayload(payload *models.ListAllServiceBindingsForServiceInstanceResponsePaged) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListAllServiceBindingsForServiceInstanceOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
