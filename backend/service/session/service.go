package session

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/syllabix/rollpay/backend/api/model"
	"github.com/syllabix/rollpay/backend/datastore/user"
	"github.com/syllabix/rollpay/backend/service/user/password"
)

type Service interface {
	Login(ctx context.Context, email, password string) (model.RollpayToken, error)
}

type service struct {
	session  *Manager
	password password.Manager
	users    user.Store
}

// Login will validate a users credentials, and if ok, create a session and return the corresponding
// access token.
func (s service) Login(ctx context.Context, email, password string) (model.RollpayToken, error) {
	user, err := s.users.GetByEmail(ctx, email)
	if err != nil {
		return failure(err)
	}

	err = s.password.Compare([]byte(user.Password), password)
	if err != nil {
		return failure(err)
	}

	token, err := s.session.Create(user.ID)
	if err != nil {
		return failure(err)
	}

	return model.RollpayToken{
		Token:      token.Value,
		Expiration: strfmt.DateTime(token.Expiration),
	}, nil
}
