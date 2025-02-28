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

	// auto clicker
	// Required: true
	AutoClicker *bool `json:"auto_clicker"`

	// click booster
	// Required: true
	ClickBooster *int64 `json:"click_booster"`

	// coin number
	// Required: true
	CoinNumber *int64 `json:"coin_number"`

	// tg id
	// Required: true
	TgID *int64 `json:"tg_id"`
}

// Validate validates this progress
func (m *Progress) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAutoClicker(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClickBooster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCoinNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTgID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Progress) validateAutoClicker(formats strfmt.Registry) error {

	if err := validate.Required("auto_clicker", "body", m.AutoClicker); err != nil {
		return err
	}

	return nil
}

func (m *Progress) validateClickBooster(formats strfmt.Registry) error {

	if err := validate.Required("click_booster", "body", m.ClickBooster); err != nil {
		return err
	}

	return nil
}

func (m *Progress) validateCoinNumber(formats strfmt.Registry) error {

	if err := validate.Required("coin_number", "body", m.CoinNumber); err != nil {
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
