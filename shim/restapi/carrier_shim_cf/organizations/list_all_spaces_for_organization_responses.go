// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/suse/carrier/shim/models"
)

// ListAllSpacesForOrganizationOKCode is the HTTP code returned for type ListAllSpacesForOrganizationOK
const ListAllSpacesForOrganizationOKCode int = 200

/*ListAllSpacesForOrganizationOK successful response

swagger:response listAllSpacesForOrganizationOK
*/
type ListAllSpacesForOrganizationOK struct {

	/*
	  In: Body
	*/
	Payload *models.ListAllSpacesForOrganizationResponsePaged `json:"body,omitempty"`
}

// NewListAllSpacesForOrganizationOK creates ListAllSpacesForOrganizationOK with default headers values
func NewListAllSpacesForOrganizationOK() *ListAllSpacesForOrganizationOK {

	return &ListAllSpacesForOrganizationOK{}
}

// WithPayload adds the payload to the list all spaces for organization o k response
func (o *ListAllSpacesForOrganizationOK) WithPayload(payload *models.ListAllSpacesForOrganizationResponsePaged) *ListAllSpacesForOrganizationOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list all spaces for organization o k response
func (o *ListAllSpacesForOrganizationOK) SetPayload(payload *models.ListAllSpacesForOrganizationResponsePaged) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListAllSpacesForOrganizationOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
