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
)

// UpdateOrganizationByIDV1MaxParseMemory sets the maximum size in bytes for
// the multipart form parser for this operation.
//
// The default value is 32 MB.
// The multipart parser stores up to this + 10MB.
var UpdateOrganizationByIDV1MaxParseMemory int64 = 32 << 20

// NewUpdateOrganizationByIDV1Params creates a new UpdateOrganizationByIDV1Params object
// with the default values initialized.
func NewUpdateOrganizationByIDV1Params() UpdateOrganizationByIDV1Params {

	var (
		// initialize parameters with default values

		acceptLanguageDefault = string("en")
		userAgentDefault      = string("test-user")
	)

	return UpdateOrganizationByIDV1Params{
		AcceptLanguage: &acceptLanguageDefault,

		UserAgent: &userAgentDefault,
	}
}

// UpdateOrganizationByIDV1Params contains all the bound params for the update organization by ID v1 operation
// typically these are obtained from a http.Request
//
// swagger:parameters UpdateOrganizationByIDV1
type UpdateOrganizationByIDV1Params struct {

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
	/*the organization logo
	  In: formData
	*/
	Logo io.ReadCloser
	/*
	  In: formData
	*/
	Name *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewUpdateOrganizationByIDV1Params() beforehand.
func (o *UpdateOrganizationByIDV1Params) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if err := r.ParseMultipartForm(UpdateOrganizationByIDV1MaxParseMemory); err != nil {
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

	rID, rhkID, _ := route.Params.GetOK("id")
	if err := o.bindID(rID, rhkID, route.Formats); err != nil {
		res = append(res, err)
	}

	logo, logoHeader, err := r.FormFile("logo")
	if err != nil && err != http.ErrMissingFile {
		res = append(res, errors.New(400, "reading file %q failed: %v", "logo", err))
	} else if err == http.ErrMissingFile {
		// no-op for missing but optional file parameter
	} else if err := o.bindLogo(logo, logoHeader); err != nil {
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
func (o *UpdateOrganizationByIDV1Params) bindAcceptLanguage(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false

	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewUpdateOrganizationByIDV1Params()
		return nil
	}
	o.AcceptLanguage = &raw

	return nil
}

// bindUserAgent binds and validates parameter UserAgent from header.
func (o *UpdateOrganizationByIDV1Params) bindUserAgent(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false

	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewUpdateOrganizationByIDV1Params()
		return nil
	}
	o.UserAgent = &raw

	return nil
}

// bindID binds and validates parameter ID from path.
func (o *UpdateOrganizationByIDV1Params) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.ID = raw

	return nil
}

// bindLogo binds file parameter Logo.
//
// The only supported validations on files are MinLength and MaxLength
func (o *UpdateOrganizationByIDV1Params) bindLogo(file multipart.File, header *multipart.FileHeader) error {
	return nil
}

// bindName binds and validates parameter Name from formData.
func (o *UpdateOrganizationByIDV1Params) bindName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.Name = &raw

	return nil
}