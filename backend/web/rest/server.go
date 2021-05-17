package rest

import (
	"net/http"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/syllabix/rollpay/backend/api/rest"
	"github.com/syllabix/rollpay/backend/api/rest/operation"
)

// Server is responsible for serving requests
type Server struct {
	http.Handler
}

// NewServer returns an http.Handler for the Toaster Rest API
func NewServer(registered Controllers) (Server, error) {

	spec, err := loads.Analyzed(rest.FlatSwaggerJSON, "")
	if err != nil {
		return Server{}, err
	}

	api := operation.NewRollpayAPI(spec)
	api.JSONConsumer = runtime.JSONConsumer()
	api.JSONProducer = runtime.JSONProducer()
	api.ServerShutdown = func() {}

	for _, r := range registered.Controllers {
		r.Register(api)
	}

	return Server{api.Serve(nil)}, nil
}
