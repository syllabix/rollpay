package user

import (
	"context"

	"github.com/syllabix/rollpay/backend/api/model"
	api "github.com/syllabix/rollpay/backend/api/rest/operation/user"
	"github.com/syllabix/rollpay/backend/common/id"
	"github.com/syllabix/rollpay/backend/datastore/user"
	"github.com/syllabix/rollpay/backend/service/user/password"
)

// func asModel(params user.CreateUserV1Params) (model.User, error) {

// 	var buf bytes.Buffer
// 	w := base64.NewEncoder(base64.RawStdEncoding, &buf)
// 	_, err := io.Copy(w, params.Avatar)
// 	if err != nil {

// 	}
// 	u.Avatar
// }

type Service interface {
	Get(ctx context.Context, id string) (model.User, error)
	Create(api.CreateUserV1Params) (model.User, error)
	Update(context.Context, model.User) (model.User, error)
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

func (s service) Update(ctx context.Context, user model.User) (model.User, error) {
	panic("not implemented") // TODO: Implement
}

func NewService(store user.Store, password password.Manager) Service {
	return service{store, password}
}
