package user

import (
	"errors"

	"github.com/syllabix/rollpay/backend/api/model"
	"github.com/syllabix/rollpay/backend/datastore/user"
)

var (
	ErrNotFound = errors.New("the requested user does not exist")
	ErrFatal    = errors.New("an unexpected system error occurred")
)

func failure(reason error) (u model.User, err error) {
	switch {
	case errors.Is(user.ErrNotFound, reason):
		err = ErrNotFound

	default:
		err = ErrFatal
	}
	return
}
