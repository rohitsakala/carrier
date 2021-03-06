// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ListEventsAssociatedWithAppSinceJanuary12014ResponseResource list events associated with app since january12014 response resource
//
// swagger:model listEventsAssociatedWithAppSinceJanuary12014ResponseResource
type ListEventsAssociatedWithAppSinceJanuary12014ResponseResource struct {

	// entity
	Entity *ListEventsAssociatedWithAppSinceJanuary12014Response `json:"entity,omitempty"`

	// metadata
	Metadata *EntityMetadata `json:"metadata,omitempty"`
}

// Validate validates this list events associated with app since january12014 response resource
func (m *ListEventsAssociatedWithAppSinceJanuary12014ResponseResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEntity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListEventsAssociatedWithAppSinceJanuary12014ResponseResource) validateEntity(formats strfmt.Registry) error {

	if swag.IsZero(m.Entity) { // not required
		return nil
	}

	if m.Entity != nil {
		if err := m.Entity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entity")
			}
			return err
		}
	}

	return nil
}

func (m *ListEventsAssociatedWithAppSinceJanuary12014ResponseResource) validateMetadata(formats strfmt.Registry) error {

	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if m.Metadata != nil {
		if err := m.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListEventsAssociatedWithAppSinceJanuary12014ResponseResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListEventsAssociatedWithAppSinceJanuary12014ResponseResource) UnmarshalBinary(b []byte) error {
	var res ListEventsAssociatedWithAppSinceJanuary12014ResponseResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
