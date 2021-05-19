package session

import (
	"errors"
	"fmt"

	"github.com/syllabix/rollpay/backend/api/model"
	"github.com/syllabix/rollpay/backend/datastore/user"
	"github.com/syllabix/rollpay/backend/service/user/password"
)

var (
	ErrFatal       = errors.New("an unexpected system error occurred")
	ErrNotFound    = errors.New("the user account does not exist")
	ErrInvalid     = errors.New("the request is invalid")
	ErrBadPassword = errors.New("the password is not correct")
)

func mapErr(reason error) (err error) {
	switch {
	case errors.Is(reason, user.ErrNotFound):
		err = ErrNotFound

	case errors.Is(reason, password.ErrSizeExceeded):
		err = fmt.Errorf("%w: %v", ErrInvalid, reason)

	case errors.Is(reason, ErrBadPassword):
		err = ErrBadPassword

	default:
		err = fmt.Errorf("%v: %w", ErrFatal, reason)
	}

	return err
}

func failure(reason error) (u model.RollpayToken, err error) {
	err = mapErr(reason)
	return
}
