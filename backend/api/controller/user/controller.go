package user

import (
	"errors"

	"github.com/go-openapi/runtime/middleware"
	"github.com/syllabix/rollpay/backend/api/model"
	"github.com/syllabix/rollpay/backend/api/rest/operation"
	"github.com/syllabix/rollpay/backend/api/rest/operation/user"
	u "github.com/syllabix/rollpay/backend/service/user"
	"github.com/syllabix/rollpay/backend/web/rest"
)

// Controller is responsible for handling media related request for the API
type Controller struct {
	srv u.Service
}

// Register controller handlers to the api
func (ctrl *Controller) Register(api *operation.RollpayAPI) {
	api.UserGetUserByIDV1Handler = user.
		GetUserByIDV1HandlerFunc(ctrl.GetUserByID)

	api.UserCreateUserV1Handler = user.
		CreateUserV1HandlerFunc(ctrl.CreateUser)
}

func (ctrl *Controller) GetUserByID(params user.GetUserByIDV1Params) middleware.Responder {
	result, err := ctrl.srv.Get(params.HTTPRequest.Context(), params.ID)
	switch {
	case err == nil:
		return user.NewGetUserByIDV1OK().
			WithPayload(&result)

	case errors.Is(err, u.ErrNotFound):
		return user.NewGetUserByIDV1NotFound().
			WithPayload(&model.StandardError{
				Message: "hmm... we were unable to find the user you were looking for",
			})

	case errors.Is(err, u.ErrInvalid):
		return user.NewCreateUserV1BadRequest().
			WithPayload(&model.StandardError{
				Message: "we are unable to process your request",
			})

	default:
		return user.NewCreateUserV1InternalServerError().
			WithPayload(&model.StandardError{
				Message: "o wow, something didn't go quite right. please grab a cool beverage and try again",
			})
	}
}

func (ctrl *Controller) CreateUser(params user.CreateUserV1Params) middleware.Responder {
	newUser, err := ctrl.srv.Create(params)
	switch {
	case err == nil:
		return user.NewCreateUserV1Created().
			WithPayload(&newUser)

	default:
		return user.NewCreateUserV1InternalServerError().
			WithPayload(&model.StandardError{
				Message: "oops, something didn't work out as expected on our end. give our systems a minute or two and try again",
			})
	}
}

// NewController intializes a new api controller for handling health endpoint
//requests
func NewController(srv u.Service) rest.Controller {
	return rest.MakeController(&Controller{
		srv: srv,
	})
}
