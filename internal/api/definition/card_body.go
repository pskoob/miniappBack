// Code generated by go-swagger; DO NOT EDIT.

package definition

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CardBody card body
//
// swagger:model CardBody
type CardBody struct {

	// auto clicker
	AutoClicker bool `json:"auto_clicker,omitempty"`

	// card id
	CardID int64 `json:"card_id,omitempty"`

	// card name
	CardName string `json:"card_name,omitempty"`

	// created at
	CreatedAt int64 `json:"created_at,omitempty"`

	// current level
	CurrentLevel int64 `json:"current_level,omitempty"`

	// updated at
	UpdatedAt int64 `json:"updated_at,omitempty"`

	// user id
	UserID int64 `json:"user_id,omitempty"`
}

// Validate validates this card body
func (m *CardBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this card body based on context it is used
func (m *CardBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CardBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CardBody) UnmarshalBinary(b []byte) error {
	var res CardBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
