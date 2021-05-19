// Code generated by go-swagger; DO NOT EDIT.

package organization

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

// NewAddOrgMembersV1Params creates a new AddOrgMembersV1Params object
// with the default values initialized.
func NewAddOrgMembersV1Params() AddOrgMembersV1Params {

	var (
		// initialize parameters with default values

		acceptLanguageDefault = string("en")
		userAgentDefault      = string("test-user")
	)

	return AddOrgMembersV1Params{
		AcceptLanguage: &acceptLanguageDefault,

		UserAgent: &userAgentDefault,
	}
}

// AddOrgMembersV1Params contains all the bound params for the add org members v1 operation
// typically these are obtained from a http.Request
//
// swagger:parameters AddOrgMembersV1
type AddOrgMembersV1Params struct {

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
	/*the id of the org
	  Required: true
	  In: path
	*/
	ID string
	/*
	  In: body
	*/
	Member AddOrgMembersV1Body
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewAddOrgMembersV1Params() beforehand.
func (o *AddOrgMembersV1Params) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
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

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body AddOrgMembersV1Body
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("member", "body", "", err))
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
				o.Member = body
			}
		}
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindAcceptLanguage binds and validates parameter AcceptLanguage from header.
func (o *AddOrgMembersV1Params) bindAcceptLanguage(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false

	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewAddOrgMembersV1Params()
		return nil
	}
	o.AcceptLanguage = &raw

	return nil
}

// bindUserAgent binds and validates parameter UserAgent from header.
func (o *AddOrgMembersV1Params) bindUserAgent(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false

	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewAddOrgMembersV1Params()
		return nil
	}
	o.UserAgent = &raw

	return nil
}

// bindID binds and validates parameter ID from path.
func (o *AddOrgMembersV1Params) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.ID = raw

	return nil
}
