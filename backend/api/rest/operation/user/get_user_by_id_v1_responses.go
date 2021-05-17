// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/syllabix/rollpay/backend/api/model"
)

// GetUserByIDV1OKCode is the HTTP code returned for type GetUserByIDV1OK
const GetUserByIDV1OKCode int = 200

/*GetUserByIDV1OK a

swagger:response getUserByIdV1OK
*/
type GetUserByIDV1OK struct {

	/*
	  In: Body
	*/
	Payload *model.User `json:"body,omitempty"`
}

// NewGetUserByIDV1OK creates GetUserByIDV1OK with default headers values
func NewGetUserByIDV1OK() *GetUserByIDV1OK {

	return &GetUserByIDV1OK{}
}

// WithPayload adds the payload to the get user by Id v1 o k response
func (o *GetUserByIDV1OK) WithPayload(payload *model.User) *GetUserByIDV1OK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user by Id v1 o k response
func (o *GetUserByIDV1OK) SetPayload(payload *model.User) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserByIDV1OK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetUserByIDV1BadRequestCode is the HTTP code returned for type GetUserByIDV1BadRequest
const GetUserByIDV1BadRequestCode int = 400

/*GetUserByIDV1BadRequest The provided request was invalid.

swagger:response getUserByIdV1BadRequest
*/
type GetUserByIDV1BadRequest struct {

	/*
	  In: Body
	*/
	Payload *model.StandardError `json:"body,omitempty"`
}

// NewGetUserByIDV1BadRequest creates GetUserByIDV1BadRequest with default headers values
func NewGetUserByIDV1BadRequest() *GetUserByIDV1BadRequest {

	return &GetUserByIDV1BadRequest{}
}

// WithPayload adds the payload to the get user by Id v1 bad request response
func (o *GetUserByIDV1BadRequest) WithPayload(payload *model.StandardError) *GetUserByIDV1BadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user by Id v1 bad request response
func (o *GetUserByIDV1BadRequest) SetPayload(payload *model.StandardError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserByIDV1BadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetUserByIDV1NotFoundCode is the HTTP code returned for type GetUserByIDV1NotFound
const GetUserByIDV1NotFoundCode int = 404

/*GetUserByIDV1NotFound The resource requested does not exist.

swagger:response getUserByIdV1NotFound
*/
type GetUserByIDV1NotFound struct {

	/*
	  In: Body
	*/
	Payload *model.StandardError `json:"body,omitempty"`
}

// NewGetUserByIDV1NotFound creates GetUserByIDV1NotFound with default headers values
func NewGetUserByIDV1NotFound() *GetUserByIDV1NotFound {

	return &GetUserByIDV1NotFound{}
}

// WithPayload adds the payload to the get user by Id v1 not found response
func (o *GetUserByIDV1NotFound) WithPayload(payload *model.StandardError) *GetUserByIDV1NotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user by Id v1 not found response
func (o *GetUserByIDV1NotFound) SetPayload(payload *model.StandardError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserByIDV1NotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetUserByIDV1InternalServerErrorCode is the HTTP code returned for type GetUserByIDV1InternalServerError
const GetUserByIDV1InternalServerErrorCode int = 500

/*GetUserByIDV1InternalServerError An unexpected system or network error occured.

swagger:response getUserByIdV1InternalServerError
*/
type GetUserByIDV1InternalServerError struct {

	/*
	  In: Body
	*/
	Payload *model.StandardError `json:"body,omitempty"`
}

// NewGetUserByIDV1InternalServerError creates GetUserByIDV1InternalServerError with default headers values
func NewGetUserByIDV1InternalServerError() *GetUserByIDV1InternalServerError {

	return &GetUserByIDV1InternalServerError{}
}

// WithPayload adds the payload to the get user by Id v1 internal server error response
func (o *GetUserByIDV1InternalServerError) WithPayload(payload *model.StandardError) *GetUserByIDV1InternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user by Id v1 internal server error response
func (o *GetUserByIDV1InternalServerError) SetPayload(payload *model.StandardError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserByIDV1InternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
