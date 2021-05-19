package session

import (
	"context"

	"github.com/syllabix/rollpay/backend/api/model"
	"github.com/syllabix/rollpay/backend/datastore/user"
	"github.com/syllabix/rollpay/backend/service/user/password"
)

type service struct {
	mngr     *Manager
	password password.Manager
	users    user.Store
}

func (s service) Login(ctx context.Context, email, password string) (model.RollpayToken, error) {
	user, err := s.users.GetByEmail(ctx, email)
	if err != nil {
		return failure(err)
	}

	err = s.password.Compare([]byte(user.Password), password)
	if err != nil {
		return failure(err)
	}

}
