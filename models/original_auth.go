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

// OriginalAuth OriginalAuth
// swagger:model OriginalAuth
type OriginalAuth struct {

	// auth plugin exists
	// Required: true
	AuthPluginExists *bool `json:"auth_plugin_exists"`

	// customer id
	// Required: true
	CustomerID *string `json:"customer_id"`

	// session id
	// Required: true
	SessionID *string `json:"session_id"`

	// session timeout
	// Required: true
	SessionTimeout *int32 `json:"session_timeout"`

	// system admin
	// Required: true
	SystemAdmin *bool `json:"system_admin"`

	// user email
	// Required: true
	UserEmail *string `json:"user_email"`

	// user id
	// Required: true
	UserID *int32 `json:"user_id"`

	// user name
	// Required: true
	UserName *string `json:"user_name"`
}

// Validate validates this original auth
func (m *OriginalAuth) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthPluginExists(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomerID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSessionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSessionTimeout(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSystemAdmin(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OriginalAuth) validateAuthPluginExists(formats strfmt.Registry) error {

	if err := validate.Required("auth_plugin_exists", "body", m.AuthPluginExists); err != nil {
		return err
	}

	return nil
}

func (m *OriginalAuth) validateCustomerID(formats strfmt.Registry) error {

	if err := validate.Required("customer_id", "body", m.CustomerID); err != nil {
		return err
	}

	return nil
}

func (m *OriginalAuth) validateSessionID(formats strfmt.Registry) error {

	if err := validate.Required("session_id", "body", m.SessionID); err != nil {
		return err
	}

	return nil
}

func (m *OriginalAuth) validateSessionTimeout(formats strfmt.Registry) error {

	if err := validate.Required("session_timeout", "body", m.SessionTimeout); err != nil {
		return err
	}

	return nil
}

func (m *OriginalAuth) validateSystemAdmin(formats strfmt.Registry) error {

	if err := validate.Required("system_admin", "body", m.SystemAdmin); err != nil {
		return err
	}

	return nil
}

func (m *OriginalAuth) validateUserEmail(formats strfmt.Registry) error {

	if err := validate.Required("user_email", "body", m.UserEmail); err != nil {
		return err
	}

	return nil
}

func (m *OriginalAuth) validateUserID(formats strfmt.Registry) error {

	if err := validate.Required("user_id", "body", m.UserID); err != nil {
		return err
	}

	return nil
}

func (m *OriginalAuth) validateUserName(formats strfmt.Registry) error {

	if err := validate.Required("user_name", "body", m.UserName); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OriginalAuth) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OriginalAuth) UnmarshalBinary(b []byte) error {
	var res OriginalAuth
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
