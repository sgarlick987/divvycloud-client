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

// AddAWSCloudAccountInstanceAssumeRoleRequest AddAWSCloudAccountInstanceAssumeRoleRequest
// swagger:model AddAWSCloudAccountInstanceAssumeRoleRequest
type AddAWSCloudAccountInstanceAssumeRoleRequest struct {

	// creation params
	// Required: true
	CreationParams *CreationParams1 `json:"creation_params"`
}

// Validate validates this add AWS cloud account instance assume role request
func (m *AddAWSCloudAccountInstanceAssumeRoleRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreationParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AddAWSCloudAccountInstanceAssumeRoleRequest) validateCreationParams(formats strfmt.Registry) error {

	if err := validate.Required("creation_params", "body", m.CreationParams); err != nil {
		return err
	}

	if m.CreationParams != nil {
		if err := m.CreationParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("creation_params")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AddAWSCloudAccountInstanceAssumeRoleRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AddAWSCloudAccountInstanceAssumeRoleRequest) UnmarshalBinary(b []byte) error {
	var res AddAWSCloudAccountInstanceAssumeRoleRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
