package user

import (
	"context"

	"github.com/syllabix/rollpay/backend/api/model"
	api "github.com/syllabix/rollpay/backend/api/rest/operation/user"
	"github.com/syllabix/rollpay/backend/datastore/user"
	"github.com/syllabix/rollpay/backend/service/user/password"
	"go.uber.org/zap"
)

// a monitor wraps and provides observability for a Service
type monitor struct {
	log *zap.Logger
	srv Service
}

func (s monitor) Get(ctx context.Context, userID string) (model.User, error) {
	result, err := s.srv.Get(ctx, userID)
	if err != nil {
		s.log.Error("failed to get a user by id",
			zap.Error(err),
			zap.String("requested-id", userID))
	}
	return result, err
}

func (s monitor) Create(params api.CreateUserV1Params) (model.User, error) {
	result, err := s.srv.Create(params)
	if err != nil {
		s.log.Error("failed to create a new user",
			zap.Error(err))
	}
	return result, err
}

func (s monitor) Update(params api.UpdateUserByIDV1Params) (model.User, error) {
	result, err := s.srv.Update(params)
	if err != nil {
		s.log.Error("failed to update a user",
			zap.Error(err))
	}
	return result, err
}

func (s monitor) Delete(ctx context.Context, userID string) error {
	err := s.srv.Delete(ctx, userID)
	if err != nil {
		s.log.Error("failed to delete a user",
			zap.Error(err))
	}
	return err
}

func NewService(store user.Store, password password.Manager, log *zap.Logger) Service {
	return monitor{
		log: log,
		srv: service{store, password},
	}
}
