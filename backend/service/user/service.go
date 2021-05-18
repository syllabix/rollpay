package user

import (
	"context"

	"github.com/syllabix/rollpay/backend/api/model"
	api "github.com/syllabix/rollpay/backend/api/rest/operation/user"
	"github.com/syllabix/rollpay/backend/common/id"
	"github.com/syllabix/rollpay/backend/common/media"
	db "github.com/syllabix/rollpay/backend/datastore/model"
	"github.com/syllabix/rollpay/backend/datastore/user"
	"github.com/syllabix/rollpay/backend/service/user/password"
)

type Service interface {
	Get(ctx context.Context, id string) (model.User, error)
	Create(api.CreateUserV1Params) (model.User, error)
	Update(api.UpdateUserByIDV1Params) (model.User, error)
	Delete(ctx context.Context, id string) error
}

type service struct {
	store    user.Store
	password password.Manager
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

func (s service) Create(params api.CreateUserV1Params) (model.User, error) {
	pwrdhash, err := s.password.GenerateHash(params.Password.String())
	if err != nil {
		return failure(err)
	}

	newUser, err := asNewUser(params, pwrdhash)
	if err != nil {
		return failure(err)
	}

	newUser, err = s.store.Create(params.HTTPRequest.Context(), newUser)
	if err != nil {
		return failure(err)
	}

	return asModel(newUser), nil
}

func (s service) Update(params api.UpdateUserByIDV1Params) (model.User, error) {
	uID, err := id.ToInternal(params.ID)
	if err != nil {
		return failure(err)
	}

	update := db.User{ID: uID}
	if params.Password != nil {
		pwrdhash, err := s.password.GenerateHash(params.Password.String())
		if err != nil {
			return failure(err)
		}
		update.Password = string(pwrdhash)
	}

	if params.Avatar != nil {
		img, err := media.Process(params.HTTPRequest.Context(), params.Avatar)
		if err != nil {
			return failure(err)
		}
		update.Avatar = img
	}

	if params.Email != nil {
		update.Email = params.Email.String()
	}

	if params.Username != nil {
		update.Username = *params.Username
	}

	update, err = s.store.Update(params.HTTPRequest.Context(), update)
	if err != nil {
		return failure(err)
	}

	return asModel(update), nil
}

func (s service) Delete(ctx context.Context, userID string) error {
	uID, err := id.ToInternal(userID)
	if err != nil {
		return mapErr(err)
	}

	err = s.store.Delete(ctx, uID)
	if err != nil {
		return mapErr(err)
	}

	return nil
}
