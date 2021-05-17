// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewPutV1UserIDParams creates a new PutV1UserIDParams object
// with the default values initialized.
func NewPutV1UserIDParams() PutV1UserIDParams {

	var (
		// initialize parameters with default values

		acceptLanguageDefault = string("en")
		userAgentDefault      = string("test-user")
	)

	return PutV1UserIDParams{
		AcceptLanguage: &acceptLanguageDefault,

		UserAgent: &userAgentDefault,
	}
}

// PutV1UserIDParams contains all the bound params for the put v1 user ID operation
// typically these are obtained from a http.Request
//
// swagger:parameters PutV1UserID
type PutV1UserIDParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*the accept language header as defined in RFC 7231, section 5.3.5 Accept-Language
	  In: header
	  Default: "en"
	*/
	AcceptLanguage *string
	/*
	  In: header
	  Default: "test-user"
	*/
	UserAgent *string
	/*the id of the user
	  Required: true
	  In: path
	*/
	ID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPutV1UserIDParams() beforehand.
func (o *PutV1UserIDParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if err := o.bindAcceptLanguage(r.Header[http.CanonicalHeaderKey("Accept-Language")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	if err := o.bindUserAgent(r.Header[http.CanonicalHeaderKey("User-Agent")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	rID, rhkID, _ := route.Params.GetOK("id")
	if err := o.bindID(rID, rhkID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindAcceptLanguage binds and validates parameter AcceptLanguage from header.
func (o *PutV1UserIDParams) bindAcceptLanguage(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false

	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewPutV1UserIDParams()
		return nil
	}
	o.AcceptLanguage = &raw

	return nil
}

// bindUserAgent binds and validates parameter UserAgent from header.
func (o *PutV1UserIDParams) bindUserAgent(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false

	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewPutV1UserIDParams()
		return nil
	}
	o.UserAgent = &raw

	return nil
}

// bindID binds and validates parameter ID from path.
func (o *PutV1UserIDParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.ID = raw

	return nil
}