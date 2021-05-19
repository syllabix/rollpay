package authorize

import (
	"errors"
	"net/http"

	apierror "github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/syllabix/rollpay/backend/api/model"
	"github.com/syllabix/rollpay/backend/api/rest/operation"
	"github.com/syllabix/rollpay/backend/api/rest/operation/authorization"
	"github.com/syllabix/rollpay/backend/service/token"
	"github.com/syllabix/rollpay/backend/web/rest"
)

// Controller is responsible for handling authorization and token
// related request for the API
type Controller struct {
	srv token.Service
}

// Register controller handlers to the api
func (ctrl *Controller) Register(api *operation.RollpayAPI) {
	api.IsAuthenticatedAuth = ctrl.Authenticate

	api.AuthorizationStartPlaidLinkV1Handler = authorization.
		StartPlaidLinkV1HandlerFunc(ctrl.StartPlaidLink)
}

// Authenticate is used to authenticate requests to the rollpay API
func (ctrl *Controller) Authenticate(token string) (*model.Principal, error) {
	// TODO: actually authenticate
	if token != "sandbox" {
		return nil, apierror.New(http.StatusUnauthorized, "please login to make this request")
	}

	return &model.Principal{}, nil
}

// StartPlaidLink handles requests to initaite a link with our payment provider Plaid
func (ctrl *Controller) StartPlaidLink(params authorization.StartPlaidLinkV1Params, session *model.Principal) middleware.Responder {
	tk, err := ctrl.srv.IssueLinkToken(params.HTTPRequest.Context(), params.User.ID)
	switch {
	case err == nil:
		return authorization.
			NewStartPlaidLinkV1Created().
			WithPayload(&tk)

	case errors.Is(err, token.ErrUnprocessable):
		return authorization.
			NewStartPlaidLinkV1BadRequest().
			WithPayload(&model.StandardError{
				// TODO: return a meaningful error message
				Message: "we could not process your request. please make sure you have an active account setup and try again",
			})

	default:
		return authorization.NewStartPlaidLinkV1InternalServerError().
			WithPayload(&model.StandardError{
				Message: "super sorry, but we were not able to link you with our payment provider. please grab a beverage of choice and try again in a few moments",
			})
	}
}

// NewController intializes a new api controller for handling health endpoint
// requests
func NewController(srv token.Service) rest.Controller {
	return rest.MakeController(&Controller{srv})
}
