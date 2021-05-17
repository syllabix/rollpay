package user

import (
	"context"

	"github.com/syllabix/rollpay/backend/api/model"
	"github.com/syllabix/rollpay/backend/datastore/user"
	"github.com/syllabix/rollpay/backend/util/id"
)

type Service interface {
	Get(ctx context.Context, id string) (model.User, error)
	Create(context.Context, model.User) (model.User, error)
	Update(context.Context, model.User) (model.User, error)
}

type service struct {
	store user.Store
}

func (s service) Get(ctx context.Context, userID string) (model.User, error) {
	uID, err := id.ToInternal(userID)
	if err != nil {
		return failure(err)
	}

	user, err := s.store.GetByID(ctx, uID)
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
