// Code generated by go-swagger; DO NOT EDIT.

package organization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"mime/multipart"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// CreateOrganizationV1MaxParseMemory sets the maximum size in bytes for
// the multipart form parser for this operation.
//
// The default value is 32 MB.
// The multipart parser stores up to this + 10MB.
var CreateOrganizationV1MaxParseMemory int64 = 32 << 20

// NewCreateOrganizationV1Params creates a new CreateOrganizationV1Params object
// with the default values initialized.
func NewCreateOrganizationV1Params() CreateOrganizationV1Params {

	var (
		// initialize parameters with default values

		acceptLanguageDefault = string("en")
		userAgentDefault      = string("test-user")
	)

	return CreateOrganizationV1Params{
		AcceptLanguage: &acceptLanguageDefault,

		UserAgent: &userAgentDefault,
	}
}

// CreateOrganizationV1Params contains all the bound params for the create organization v1 operation
// typically these are obtained from a http.Request
//
// swagger:parameters CreateOrganizationV1
type CreateOrganizationV1Params struct {

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
	/*the organization logo
	  Required: true
	  In: formData
	*/
	Logo io.ReadCloser
	/*
	  Required: true
	  In: formData
	*/
	Name string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewCreateOrganizationV1Params() beforehand.
func (o *CreateOrganizationV1Params) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if err := r.ParseMultipartForm(CreateOrganizationV1MaxParseMemory); err != nil {
		if err != http.ErrNotMultipart {
			return errors.New(400, "%v", err)
		} else if err := r.ParseForm(); err != nil {
			return errors.New(400, "%v", err)
		}
	}
	fds := runtime.Values(r.Form)

	if err := o.bindAcceptLanguage(r.Header[http.CanonicalHeaderKey("Accept-Language")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	if err := o.bindUserAgent(r.Header[http.CanonicalHeaderKey("User-Agent")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	logo, logoHeader, err := r.FormFile("logo")
	if err != nil {
		res = append(res, errors.New(400, "reading file %q failed: %v", "logo", err))
	} else if err := o.bindLogo(logo, logoHeader); err != nil {
		// Required: true
		res = append(res, err)
	} else {
		o.Logo = &runtime.File{Data: logo, Header: logoHeader}
	}

	fdName, fdhkName, _ := fds.GetOK("name")
	if err := o.bindName(fdName, fdhkName, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindAcceptLanguage binds and validates parameter AcceptLanguage from header.
func (o *CreateOrganizationV1Params) bindAcceptLanguage(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false

	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewCreateOrganizationV1Params()
		return nil
	}
	o.AcceptLanguage = &raw

	return nil
}

// bindUserAgent binds and validates parameter UserAgent from header.
func (o *CreateOrganizationV1Params) bindUserAgent(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false

	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewCreateOrganizationV1Params()
		return nil
	}
	o.UserAgent = &raw

	return nil
}

// bindLogo binds file parameter Logo.
//
// The only supported validations on files are MinLength and MaxLength
func (o *CreateOrganizationV1Params) bindLogo(file multipart.File, header *multipart.FileHeader) error {
	return nil
}

// bindName binds and validates parameter Name from formData.
func (o *CreateOrganizationV1Params) bindName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("name", "formData", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true

	if err := validate.RequiredString("name", "formData", raw); err != nil {
		return err
	}
	o.Name = raw

	return nil
}
