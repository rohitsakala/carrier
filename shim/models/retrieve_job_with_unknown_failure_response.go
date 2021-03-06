// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RetrieveJobWithUnknownFailureResponse retrieve job with unknown failure response
//
// swagger:model retrieveJobWithUnknownFailureResponse
type RetrieveJobWithUnknownFailureResponse struct {

	// The error
	Error string `json:"error,omitempty"`

	// The error Details
	ErrorDetails GenericObject `json:"error_details,omitempty"`

	// The guid
	GUID string `json:"guid,omitempty"`

	// The status
	Status string `json:"status,omitempty"`
}

// Validate validates this retrieve job with unknown failure response
func (m *RetrieveJobWithUnknownFailureResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrorDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RetrieveJobWithUnknownFailureResponse) validateErrorDetails(formats strfmt.Registry) error {

	if swag.IsZero(m.ErrorDetails) { // not required
		return nil
	}

	if err := m.ErrorDetails.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("error_details")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RetrieveJobWithUnknownFailureResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RetrieveJobWithUnknownFailureResponse) UnmarshalBinary(b []byte) error {
	var res RetrieveJobWithUnknownFailureResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
