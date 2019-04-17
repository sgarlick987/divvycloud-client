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

// QueryResourceRequest QueryResourceRequest
// swagger:model QueryResourceRequest
type QueryResourceRequest struct {

	// filters
	// Required: true
	Filters []string `json:"filters"`

	// limit
	// Required: true
	Limit *int32 `json:"limit"`

	// offset
	// Required: true
	Offset *int32 `json:"offset"`

	// order by
	// Required: true
	OrderBy *string `json:"order_by"`

	// scopes
	// Required: true
	Scopes []string `json:"scopes"`

	// selected resource type
	// Required: true
	SelectedResourceType *string `json:"selected_resource_type"`
}

// Validate validates this query resource request
func (m *QueryResourceRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFilters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLimit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOffset(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrderBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScopes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelectedResourceType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QueryResourceRequest) validateFilters(formats strfmt.Registry) error {

	if err := validate.Required("filters", "body", m.Filters); err != nil {
		return err
	}

	return nil
}

func (m *QueryResourceRequest) validateLimit(formats strfmt.Registry) error {

	if err := validate.Required("limit", "body", m.Limit); err != nil {
		return err
	}

	return nil
}

func (m *QueryResourceRequest) validateOffset(formats strfmt.Registry) error {

	if err := validate.Required("offset", "body", m.Offset); err != nil {
		return err
	}

	return nil
}

func (m *QueryResourceRequest) validateOrderBy(formats strfmt.Registry) error {

	if err := validate.Required("order_by", "body", m.OrderBy); err != nil {
		return err
	}

	return nil
}

func (m *QueryResourceRequest) validateScopes(formats strfmt.Registry) error {

	if err := validate.Required("scopes", "body", m.Scopes); err != nil {
		return err
	}

	return nil
}

func (m *QueryResourceRequest) validateSelectedResourceType(formats strfmt.Registry) error {

	if err := validate.Required("selected_resource_type", "body", m.SelectedResourceType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *QueryResourceRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QueryResourceRequest) UnmarshalBinary(b []byte) error {
	var res QueryResourceRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
