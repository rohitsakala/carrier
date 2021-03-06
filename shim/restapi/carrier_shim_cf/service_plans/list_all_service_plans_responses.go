// Code generated by go-swagger; DO NOT EDIT.

package service_plans

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/suse/carrier/shim/models"
)

// ListAllServicePlansOKCode is the HTTP code returned for type ListAllServicePlansOK
const ListAllServicePlansOKCode int = 200

/*ListAllServicePlansOK successful response

swagger:response listAllServicePlansOK
*/
type ListAllServicePlansOK struct {

	/*
	  In: Body
	*/
	Payload *models.ListAllServicePlansResponsePaged `json:"body,omitempty"`
}

// NewListAllServicePlansOK creates ListAllServicePlansOK with default headers values
func NewListAllServicePlansOK() *ListAllServicePlansOK {

	return &ListAllServicePlansOK{}
}

// WithPayload adds the payload to the list all service plans o k response
func (o *ListAllServicePlansOK) WithPayload(payload *models.ListAllServicePlansResponsePaged) *ListAllServicePlansOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list all service plans o k response
func (o *ListAllServicePlansOK) SetPayload(payload *models.ListAllServicePlansResponsePaged) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListAllServicePlansOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
