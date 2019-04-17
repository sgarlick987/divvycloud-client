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

// CreationParams CreationParams
// swagger:model CreationParams
type CreationParams struct {

	// account number
	// Required: true
	AccountNumber *string `json:"account_number"`

	// api key
	// Required: true
	APIKey *string `json:"api_key"`

	// authentication type
	// Required: true
	AuthenticationType *string `json:"authentication_type"`

	// cloud type
	// Required: true
	CloudType *string `json:"cloud_type"`

	// name
	// Required: true
	Name *string `json:"name"`

	// secret key
	// Required: true
	SecretKey *string `json:"secret_key"`
}

// Validate validates this creation params
func (m *CreationParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAPIKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthenticationType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCloudType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecretKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreationParams) validateAccountNumber(formats strfmt.Registry) error {

	if err := validate.Required("account_number", "body", m.AccountNumber); err != nil {
		return err
	}

	return nil
}

func (m *CreationParams) validateAPIKey(formats strfmt.Registry) error {

	if err := validate.Required("api_key", "body", m.APIKey); err != nil {
		return err
	}

	return nil
}

func (m *CreationParams) validateAuthenticationType(formats strfmt.Registry) error {

	if err := validate.Required("authentication_type", "body", m.AuthenticationType); err != nil {
		return err
	}

	return nil
}

func (m *CreationParams) validateCloudType(formats strfmt.Registry) error {

	if err := validate.Required("cloud_type", "body", m.CloudType); err != nil {
		return err
	}

	return nil
}

func (m *CreationParams) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *CreationParams) validateSecretKey(formats strfmt.Registry) error {

	if err := validate.Required("secret_key", "body", m.SecretKey); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreationParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreationParams) UnmarshalBinary(b []byte) error {
	var res CreationParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
