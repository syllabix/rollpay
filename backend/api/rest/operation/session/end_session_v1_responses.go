// Code generated by go-swagger; DO NOT EDIT.

package session

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/syllabix/rollpay/backend/api/model"
)

// EndSessionV1OKCode is the HTTP code returned for type EndSessionV1OK
const EndSessionV1OKCode int = 200

/*EndSessionV1OK session has been terminated

swagger:response endSessionV1OK
*/
type EndSessionV1OK struct {

	/*
	  In: Body
	*/
	Payload *model.StandardResponse `json:"body,omitempty"`
}

// NewEndSessionV1OK creates EndSessionV1OK with default headers values
func NewEndSessionV1OK() *EndSessionV1OK {

	return &EndSessionV1OK{}
}

// WithPayload adds the payload to the end session v1 o k response
func (o *EndSessionV1OK) WithPayload(payload *model.StandardResponse) *EndSessionV1OK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the end session v1 o k response
func (o *EndSessionV1OK) SetPayload(payload *model.StandardResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *EndSessionV1OK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// EndSessionV1BadRequestCode is the HTTP code returned for type EndSessionV1BadRequest
const EndSessionV1BadRequestCode int = 400

/*EndSessionV1BadRequest The provided request was invalid.

swagger:response endSessionV1BadRequest
*/
type EndSessionV1BadRequest struct {

	/*
	  In: Body
	*/
	Payload *model.StandardError `json:"body,omitempty"`
}

// NewEndSessionV1BadRequest creates EndSessionV1BadRequest with default headers values
func NewEndSessionV1BadRequest() *EndSessionV1BadRequest {

	return &EndSessionV1BadRequest{}
}

// WithPayload adds the payload to the end session v1 bad request response
func (o *EndSessionV1BadRequest) WithPayload(payload *model.StandardError) *EndSessionV1BadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the end session v1 bad request response
func (o *EndSessionV1BadRequest) SetPayload(payload *model.StandardError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *EndSessionV1BadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// EndSessionV1UnauthorizedCode is the HTTP code returned for type EndSessionV1Unauthorized
const EndSessionV1UnauthorizedCode int = 401

/*EndSessionV1Unauthorized The requested resource requires authentication.

swagger:response endSessionV1Unauthorized
*/
type EndSessionV1Unauthorized struct {

	/*
	  In: Body
	*/
	Payload *model.StandardError `json:"body,omitempty"`
}

// NewEndSessionV1Unauthorized creates EndSessionV1Unauthorized with default headers values
func NewEndSessionV1Unauthorized() *EndSessionV1Unauthorized {

	return &EndSessionV1Unauthorized{}
}

// WithPayload adds the payload to the end session v1 unauthorized response
func (o *EndSessionV1Unauthorized) WithPayload(payload *model.StandardError) *EndSessionV1Unauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the end session v1 unauthorized response
func (o *EndSessionV1Unauthorized) SetPayload(payload *model.StandardError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *EndSessionV1Unauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// EndSessionV1InternalServerErrorCode is the HTTP code returned for type EndSessionV1InternalServerError
const EndSessionV1InternalServerErrorCode int = 500

/*EndSessionV1InternalServerError An unexpected system or network error occured.

swagger:response endSessionV1InternalServerError
*/
type EndSessionV1InternalServerError struct {

	/*
	  In: Body
	*/
	Payload *model.StandardError `json:"body,omitempty"`
}

// NewEndSessionV1InternalServerError creates EndSessionV1InternalServerError with default headers values
func NewEndSessionV1InternalServerError() *EndSessionV1InternalServerError {

	return &EndSessionV1InternalServerError{}
}

// WithPayload adds the payload to the end session v1 internal server error response
func (o *EndSessionV1InternalServerError) WithPayload(payload *model.StandardError) *EndSessionV1InternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the end session v1 internal server error response
func (o *EndSessionV1InternalServerError) SetPayload(payload *model.StandardError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *EndSessionV1InternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
