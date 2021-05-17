package user

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/syllabix/rollpay/backend/datastore/model"
)

var (
	ErrNotFound = errors.New("the requested user was not found")
	ErrFatal    = errors.New("an unexpected error occurred while accessing the database")
)

func failure(reason error) (user model.User, err error) {
	switch {
	case errors.Is(reason, sql.ErrNoRows):
		err = ErrNotFound
	default:
		err = fmt.Errorf("%w: %v", ErrFatal, reason)
	}

	return
}
