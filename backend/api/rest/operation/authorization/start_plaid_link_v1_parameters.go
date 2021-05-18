// Code generated by go-swagger; DO NOT EDIT.

package authorization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewStartPlaidLinkV1Params creates a new StartPlaidLinkV1Params object
// with the default values initialized.
func NewStartPlaidLinkV1Params() StartPlaidLinkV1Params {

	var (
		// initialize parameters with default values

		acceptLanguageDefault = string("en")
		userAgentDefault      = string("test-user")
	)

	return StartPlaidLinkV1Params{
		AcceptLanguage: &acceptLanguageDefault,

		UserAgent: &userAgentDefault,
	}
}

// StartPlaidLinkV1Params contains all the bound params for the start plaid link v1 operation
// typically these are obtained from a http.Request
//
// swagger:parameters StartPlaidLinkV1
type StartPlaidLinkV1Params struct {

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
	/*
	  In: body
	*/
	User StartPlaidLinkV1Body
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewStartPlaidLinkV1Params() beforehand.
func (o *StartPlaidLinkV1Params) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if err := o.bindAcceptLanguage(r.Header[http.CanonicalHeaderKey("Accept-Language")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	if err := o.bindUserAgent(r.Header[http.CanonicalHeaderKey("User-Agent")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body StartPlaidLinkV1Body
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("user", "body", "", err))
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(context.Background())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.User = body
			}
		}
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindAcceptLanguage binds and validates parameter AcceptLanguage from header.
func (o *StartPlaidLinkV1Params) bindAcceptLanguage(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false

	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewStartPlaidLinkV1Params()
		return nil
	}
	o.AcceptLanguage = &raw

	return nil
}

// bindUserAgent binds and validates parameter UserAgent from header.
func (o *StartPlaidLinkV1Params) bindUserAgent(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false

	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewStartPlaidLinkV1Params()
		return nil
	}
	o.UserAgent = &raw

	return nil
}