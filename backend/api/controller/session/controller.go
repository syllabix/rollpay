package session

import (
	"errors"

	"github.com/go-openapi/runtime/middleware"
	"github.com/syllabix/rollpay/backend/api/model"
	"github.com/syllabix/rollpay/backend/api/rest/operation"
	"github.com/syllabix/rollpay/backend/api/rest/operation/session"
	s "github.com/syllabix/rollpay/backend/service/session"
	"github.com/syllabix/rollpay/backend/web/rest"
)

// Controller is responsible for handling login/logout requests for the API
type Controller struct {
	srv s.Service
}

// Register controller handlers to the api
func (ctrl *Controller) Register(api *operation.RollpayAPI) {
	api.SessionStartSessionV1Handler = session.
		StartSessionV1HandlerFunc(ctrl.Login)

	api.SessionEndSessionV1Handler = session.
		EndSessionV1HandlerFunc(ctrl.Logout)
}

// HealthCheck handles health check requests to the toaster api
func (ctrl *Controller) Login(params session.StartSessionV1Params) middleware.Responder {
	tk, err := ctrl.srv.Login(
		params.HTTPRequest.Context(),
		params.Credentials.Email.String(),
		params.Credentials.Password.String())

	switch {
	case err == nil:
		return session.
			NewStartSessionV1Created().
			WithPayload(&tk)

	case errors.Is(err, s.ErrBadPassword),
		errors.Is(err, s.ErrNotFound):
		return session.
			NewStartSessionV1Unauthorized().
			WithPayload(&model.StandardError{
				Message: "the email and password combination ",
			})

	default:
		return session.
			NewStartSessionV1InternalServerError().
			WithPayload(&model.StandardError{
				Message: "well this is embarrassing, something happened while trying to log you in",
			})
	}

}

// HealthCheck handles health check requests to the toaster api
func (ctrl *Controller) Logout(params session.EndSessionV1Params, sec *model.Principal) middleware.Responder {
	return nil
}

// NewController intializes a new api controller for handling health endpoint
// requests
func NewController(srv s.Service) rest.Controller {
	return rest.MakeController(&Controller{srv})
}
