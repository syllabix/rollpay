package user

import (
	"errors"
	"fmt"

	"github.com/syllabix/rollpay/backend/api/model"
	"github.com/syllabix/rollpay/backend/common/id"
	"github.com/syllabix/rollpay/backend/datastore/user"
	"github.com/syllabix/rollpay/backend/service/user/password"
)

var (
	ErrFatal         = errors.New("an unexpected system error occurred")
	ErrNotFound      = errors.New("the requested user does not exist")
	ErrInvalid       = errors.New("the request is invalid")
	ErrEmailReserved = errors.New("the email used has already reserved")
)

func mapErr(reason error) (err error) {
	switch {
	case errors.Is(reason, user.ErrNotFound):
		err = ErrNotFound

	case errors.Is(reason, user.ErrEmailTaken):
		err = ErrEmailReserved

	case errors.Is(reason, id.ErrInvalid),
		errors.Is(reason, password.ErrSizeExceeded):
		err = fmt.Errorf("%w: %v", ErrInvalid, reason)

	default:
		err = fmt.Errorf("%v: %w", ErrFatal, err)
	}

	return err
}

func failure(reason error) (u model.User, err error) {
	err = mapErr(reason)
	return
}
