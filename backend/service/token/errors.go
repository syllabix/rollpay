package token

import (
	"errors"
	"fmt"

	"github.com/syllabix/rollpay/backend/api/model"
	"github.com/syllabix/rollpay/backend/common/id"
	"github.com/syllabix/rollpay/backend/datastore/user"
)

var (
	ErrFatal         = errors.New("a fatal error occurred while processing your token request")
	ErrUnprocessable = errors.New("the token request cannot be processed")
)

func mapErr(reason error) (err error) {
	switch {
	case errors.Is(reason, id.ErrInvalid),
		errors.Is(reason, user.ErrNotFound):
		err = ErrUnprocessable

	default:
		err = fmt.Errorf("%w: %v", ErrFatal, reason)
	}

	return err
}

func failure(reason error) (tk model.LinkToken, err error) {
	err = mapErr(reason)
	return
}
