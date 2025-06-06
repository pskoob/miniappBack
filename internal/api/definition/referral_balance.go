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

// ReferralBalance referral balance
//
// swagger:model ReferralBalance
type ReferralBalance struct {

	// referral balance
	// Required: true
	ReferralBalance *int64 `json:"referral_balance"`
}

// Validate validates this referral balance
func (m *ReferralBalance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReferralBalance(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReferralBalance) validateReferralBalance(formats strfmt.Registry) error {

	if err := validate.Required("referral_balance", "body", m.ReferralBalance); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this referral balance based on context it is used
func (m *ReferralBalance) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ReferralBalance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReferralBalance) UnmarshalBinary(b []byte) error {
	var res ReferralBalance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
