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

// Progress progress
//
// swagger:model Progress
type Progress struct {

	// click count
	// Required: true
	ClickCount *int64 `json:"click_count"`

	// has auto clicker
	// Required: true
	HasAutoClicker *bool `json:"has_auto_clicker"`

	// tg id
	// Required: true
	TgID *int64 `json:"tg_id"`

	// upgrade energy
	// Required: true
	UpgradeEnergy *int64 `json:"upgrade_energy"`

	// upgrade level
	// Required: true
	UpgradeLevel *int64 `json:"upgrade_level"`
}

// Validate validates this progress
func (m *Progress) Validate(formats strfmt.Registry) error {
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

	if err := m.validateUpgradeEnergy(formats); err != nil {
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

func (m *Progress) validateClickCount(formats strfmt.Registry) error {

	if err := validate.Required("click_count", "body", m.ClickCount); err != nil {
		return err
	}

	return nil
}

func (m *Progress) validateHasAutoClicker(formats strfmt.Registry) error {

	if err := validate.Required("has_auto_clicker", "body", m.HasAutoClicker); err != nil {
		return err
	}

	return nil
}

func (m *Progress) validateTgID(formats strfmt.Registry) error {

	if err := validate.Required("tg_id", "body", m.TgID); err != nil {
		return err
	}

	return nil
}

func (m *Progress) validateUpgradeEnergy(formats strfmt.Registry) error {

	if err := validate.Required("upgrade_energy", "body", m.UpgradeEnergy); err != nil {
		return err
	}

	return nil
}

func (m *Progress) validateUpgradeLevel(formats strfmt.Registry) error {

	if err := validate.Required("upgrade_level", "body", m.UpgradeLevel); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this progress based on context it is used
func (m *Progress) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Progress) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Progress) UnmarshalBinary(b []byte) error {
	var res Progress
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
