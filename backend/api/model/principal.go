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

// Principal principal
//
// swagger:model Principal
type Principal struct {

	// session Id
	// Read Only: true
	SessionID int64 `json:"sessionId,omitempty"`

	// user Id
	// Read Only: true
	UserID int64 `json:"userId,omitempty"`
}

// Validate validates this principal
func (m *Principal) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this principal based on the context it is used
func (m *Principal) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSessionID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUserID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Principal) contextValidateSessionID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "sessionId", "body", int64(m.SessionID)); err != nil {
		return err
	}

	return nil
}

func (m *Principal) contextValidateUserID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "userId", "body", int64(m.UserID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Principal) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Principal) UnmarshalBinary(b []byte) error {
	var res Principal
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}