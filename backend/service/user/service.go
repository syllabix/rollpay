package user

import (
	"context"

	"github.com/syllabix/rollpay/backend/api/model"
	"github.com/syllabix/rollpay/backend/datastore/user"
)

type Service interface {
	Get(ctx context.Context, id int64) (model.User, error)
	Create(context.Context, model.User) (model.User, error)
	Update(context.Context, model.User) (model.User, error)
}

type service struct {
	store user.Store
}

func (s service) Get(ctx context.Context, id int64) (model.User, error) {
	user, err := s.store.GetByID(ctx, id)
	if err != nil {
		return failure(err)
	}

	return asModel(user), nil
}

func (s service) Create(ctx context.Context, user model.User) (model.User, error) {
	panic("not implemented") // TODO: Implement
}

func (s service) Update(ctx context.Context, user model.User) (model.User, error) {
	panic("not implemented") // TODO: Implement
}

func NewService(store user.Store) Service {
	return service{store}
}
