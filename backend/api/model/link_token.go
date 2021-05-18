// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// LinkToken link token
//
// swagger:model LinkToken
type LinkToken struct {

	// token
	// Read Only: true
	Token string `json:"token,omitempty"`
}

// Validate validates this link token
func (m *LinkToken) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this link token based on the context it is used
func (m *LinkToken) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateToken(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LinkToken) contextValidateToken(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "token", "body", string(m.Token)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LinkToken) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LinkToken) UnmarshalBinary(b []byte) error {
	var res LinkToken
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
