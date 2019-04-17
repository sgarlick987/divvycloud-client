// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Details1 Details1
// swagger:model Details1
type Details1 struct {

	// resource type
	// Required: true
	ResourceType *string `json:"resource_type"`

	// storagecontainer
	// Required: true
	Storagecontainer *Storagecontainer `json:"storagecontainer"`
}

// Validate validates this details1
func (m *Details1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResourceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStoragecontainer(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Details1) validateResourceType(formats strfmt.Registry) error {

	if err := validate.Required("resource_type", "body", m.ResourceType); err != nil {
		return err
	}

	return nil
}

func (m *Details1) validateStoragecontainer(formats strfmt.Registry) error {

	if err := validate.Required("storagecontainer", "body", m.Storagecontainer); err != nil {
		return err
	}

	if m.Storagecontainer != nil {
		if err := m.Storagecontainer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("storagecontainer")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Details1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Details1) UnmarshalBinary(b []byte) error {
	var res Details1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
