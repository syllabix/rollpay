// Code generated by go-swagger; DO NOT EDIT.

package health

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
)

// NewCheckV1Params creates a new CheckV1Params object
//
// There are no default values defined in the spec.
func NewCheckV1Params() CheckV1Params {

	return CheckV1Params{}
}

// CheckV1Params contains all the bound params for the check v1 operation
// typically these are obtained from a http.Request
//
// swagger:parameters CheckV1
type CheckV1Params struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewCheckV1Params() beforehand.
func (o *CheckV1Params) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}