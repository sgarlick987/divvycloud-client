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

// Properties Properties
// swagger:model Properties
type Properties struct {

	// property list
	// Required: true
	PropertyList []string `json:"property_list"`
}

// Validate validates this properties
func (m *Properties) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePropertyList(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Properties) validatePropertyList(formats strfmt.Registry) error {

	if err := validate.Required("property_list", "body", m.PropertyList); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Properties) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Properties) UnmarshalBinary(b []byte) error {
	var res Properties
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
