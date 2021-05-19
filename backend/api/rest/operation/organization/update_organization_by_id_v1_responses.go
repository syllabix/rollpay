// Code generated by go-swagger; DO NOT EDIT.

package organization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/syllabix/rollpay/backend/api/model"
)

// UpdateOrganizationByIDV1OKCode is the HTTP code returned for type UpdateOrganizationByIDV1OK
const UpdateOrganizationByIDV1OKCode int = 200

/*UpdateOrganizationByIDV1OK a successfully updated organization

swagger:response updateOrganizationByIdV1OK
*/
type UpdateOrganizationByIDV1OK struct {

	/*
	  In: Body
	*/
	Payload *model.Organization `json:"body,omitempty"`
}

// NewUpdateOrganizationByIDV1OK creates UpdateOrganizationByIDV1OK with default headers values
func NewUpdateOrganizationByIDV1OK() *UpdateOrganizationByIDV1OK {

	return &UpdateOrganizationByIDV1OK{}
}

// WithPayload adds the payload to the update organization by Id v1 o k response
func (o *UpdateOrganizationByIDV1OK) WithPayload(payload *model.Organization) *UpdateOrganizationByIDV1OK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update organization by Id v1 o k response
func (o *UpdateOrganizationByIDV1OK) SetPayload(payload *model.Organization) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateOrganizationByIDV1OK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateOrganizationByIDV1BadRequestCode is the HTTP code returned for type UpdateOrganizationByIDV1BadRequest
const UpdateOrganizationByIDV1BadRequestCode int = 400

/*UpdateOrganizationByIDV1BadRequest The provided request was invalid.

swagger:response updateOrganizationByIdV1BadRequest
*/
type UpdateOrganizationByIDV1BadRequest struct {

	/*
	  In: Body
	*/
	Payload *model.StandardError `json:"body,omitempty"`
}

// NewUpdateOrganizationByIDV1BadRequest creates UpdateOrganizationByIDV1BadRequest with default headers values
func NewUpdateOrganizationByIDV1BadRequest() *UpdateOrganizationByIDV1BadRequest {

	return &UpdateOrganizationByIDV1BadRequest{}
}

// WithPayload adds the payload to the update organization by Id v1 bad request response
func (o *UpdateOrganizationByIDV1BadRequest) WithPayload(payload *model.StandardError) *UpdateOrganizationByIDV1BadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update organization by Id v1 bad request response
func (o *UpdateOrganizationByIDV1BadRequest) SetPayload(payload *model.StandardError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateOrganizationByIDV1BadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateOrganizationByIDV1NotFoundCode is the HTTP code returned for type UpdateOrganizationByIDV1NotFound
const UpdateOrganizationByIDV1NotFoundCode int = 404

/*UpdateOrganizationByIDV1NotFound The resource requested does not exist.

swagger:response updateOrganizationByIdV1NotFound
*/
type UpdateOrganizationByIDV1NotFound struct {

	/*
	  In: Body
	*/
	Payload *model.StandardError `json:"body,omitempty"`
}

// NewUpdateOrganizationByIDV1NotFound creates UpdateOrganizationByIDV1NotFound with default headers values
func NewUpdateOrganizationByIDV1NotFound() *UpdateOrganizationByIDV1NotFound {

	return &UpdateOrganizationByIDV1NotFound{}
}

// WithPayload adds the payload to the update organization by Id v1 not found response
func (o *UpdateOrganizationByIDV1NotFound) WithPayload(payload *model.StandardError) *UpdateOrganizationByIDV1NotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update organization by Id v1 not found response
func (o *UpdateOrganizationByIDV1NotFound) SetPayload(payload *model.StandardError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateOrganizationByIDV1NotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateOrganizationByIDV1ConflictCode is the HTTP code returned for type UpdateOrganizationByIDV1Conflict
const UpdateOrganizationByIDV1ConflictCode int = 409

/*UpdateOrganizationByIDV1Conflict A conflict with an existing resource or process occured.

swagger:response updateOrganizationByIdV1Conflict
*/
type UpdateOrganizationByIDV1Conflict struct {

	/*
	  In: Body
	*/
	Payload *model.StandardError `json:"body,omitempty"`
}

// NewUpdateOrganizationByIDV1Conflict creates UpdateOrganizationByIDV1Conflict with default headers values
func NewUpdateOrganizationByIDV1Conflict() *UpdateOrganizationByIDV1Conflict {

	return &UpdateOrganizationByIDV1Conflict{}
}

// WithPayload adds the payload to the update organization by Id v1 conflict response
func (o *UpdateOrganizationByIDV1Conflict) WithPayload(payload *model.StandardError) *UpdateOrganizationByIDV1Conflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update organization by Id v1 conflict response
func (o *UpdateOrganizationByIDV1Conflict) SetPayload(payload *model.StandardError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateOrganizationByIDV1Conflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateOrganizationByIDV1InternalServerErrorCode is the HTTP code returned for type UpdateOrganizationByIDV1InternalServerError
const UpdateOrganizationByIDV1InternalServerErrorCode int = 500

/*UpdateOrganizationByIDV1InternalServerError An unexpected system or network error occured.

swagger:response updateOrganizationByIdV1InternalServerError
*/
type UpdateOrganizationByIDV1InternalServerError struct {

	/*
	  In: Body
	*/
	Payload *model.StandardError `json:"body,omitempty"`
}

// NewUpdateOrganizationByIDV1InternalServerError creates UpdateOrganizationByIDV1InternalServerError with default headers values
func NewUpdateOrganizationByIDV1InternalServerError() *UpdateOrganizationByIDV1InternalServerError {

	return &UpdateOrganizationByIDV1InternalServerError{}
}

// WithPayload adds the payload to the update organization by Id v1 internal server error response
func (o *UpdateOrganizationByIDV1InternalServerError) WithPayload(payload *model.StandardError) *UpdateOrganizationByIDV1InternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update organization by Id v1 internal server error response
func (o *UpdateOrganizationByIDV1InternalServerError) SetPayload(payload *model.StandardError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateOrganizationByIDV1InternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}