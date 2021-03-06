// Code generated by go-swagger; DO NOT EDIT.

package session

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// StartSessionV1HandlerFunc turns a function with the right signature into a start session v1 handler
type StartSessionV1HandlerFunc func(StartSessionV1Params) middleware.Responder

// Handle executing the request and returning a response
func (fn StartSessionV1HandlerFunc) Handle(params StartSessionV1Params) middleware.Responder {
	return fn(params)
}

// StartSessionV1Handler interface for that can handle valid start session v1 params
type StartSessionV1Handler interface {
	Handle(StartSessionV1Params) middleware.Responder
}

// NewStartSessionV1 creates a new http.Handler for the start session v1 operation
func NewStartSessionV1(ctx *middleware.Context, handler StartSessionV1Handler) *StartSessionV1 {
	return &StartSessionV1{Context: ctx, Handler: handler}
}

/* StartSessionV1 swagger:route POST /v1/login Session startSessionV1

attempt to login and receive an auth token for the service

*/
type StartSessionV1 struct {
	Context *middleware.Context
	Handler StartSessionV1Handler
}

func (o *StartSessionV1) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewStartSessionV1Params()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
