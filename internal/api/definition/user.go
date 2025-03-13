// Code generated by go-swagger; DO NOT EDIT.

package definition

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// User user
//
// swagger:model User
type User struct {

	// click count
	// Required: true
	ClickCount *int64 `json:"clickCount"`

	// has auto clicker
	// Required: true
	HasAutoClicker *bool `json:"hasAutoClicker"`

	// tg id
	// Required: true
	TgID *int64 `json:"tg_id"`

	// upgrade level
	// Required: true
	UpgradeLevel *int64 `json:"upgradeLevel"`
	
}

// Validate validates this user
func (m *User) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClickCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHasAutoClicker(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTgID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpgradeLevel(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *User) validateClickCount(formats strfmt.Registry) error {

	if err := validate.Required("clickCount", "body", m.ClickCount); err != nil {
		return err
	}

	return nil
}

func (m *User) validateHasAutoClicker(formats strfmt.Registry) error {

	if err := validate.Required("hasAutoClicker", "body", m.HasAutoClicker); err != nil {
		return err
	}

	return nil
}

func (m *User) validateTgID(formats strfmt.Registry) error {

	if err := validate.Required("tg_id", "body", m.TgID); err != nil {
		return err
	}

	return nil
}

func (m *User) validateUpgradeLevel(formats strfmt.Registry) error {

	if err := validate.Required("upgradeLevel", "body", m.UpgradeLevel); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this user based on context it is used
func (m *User) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *User) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *User) UnmarshalBinary(b []byte) error {
	var res User
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
