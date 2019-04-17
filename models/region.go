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

// Region Region
// swagger:model Region
type Region struct {

	// harvest rate multiplier
	// Required: true
	HarvestRateMultiplier *float64 `json:"harvest_rate_multiplier"`

	// id
	// Required: true
	ID *string `json:"id"`

	// name
	// Required: true
	Name *string `json:"name"`

	// resource id
	// Required: true
	ResourceID *string `json:"resource_id"`

	// status
	// Required: true
	Status *string `json:"status"`
}

// Validate validates this region
func (m *Region) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHarvestRateMultiplier(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Region) validateHarvestRateMultiplier(formats strfmt.Registry) error {

	if err := validate.Required("harvest_rate_multiplier", "body", m.HarvestRateMultiplier); err != nil {
		return err
	}

	return nil
}

func (m *Region) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *Region) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Region) validateResourceID(formats strfmt.Registry) error {

	if err := validate.Required("resource_id", "body", m.ResourceID); err != nil {
		return err
	}

	return nil
}

func (m *Region) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Region) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Region) UnmarshalBinary(b []byte) error {
	var res Region
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
