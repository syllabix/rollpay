package user

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/lib/pq"
	"github.com/syllabix/rollpay/backend/datastore/model"
)

var (
	ErrNotFound   = errors.New("the requested user was not found")
	ErrFatal      = errors.New("an unexpected error occurred while accessing the database")
	ErrEmailTaken = errors.New("the email being used to create a new account is already in use")
)

func failure(reason error) (user model.User, err error) {
	var sqlErr *pq.Error
	switch {
	case errors.Is(reason, sql.ErrNoRows):
		err = ErrNotFound

	case errors.As(reason, &sqlErr):
		if sqlErr.Code.Name() == "unique_violation" {
			err = ErrEmailTaken
			break
		}
		// if it is not an expected error, then use defaul case
		fallthrough

	default:
		err = fmt.Errorf("%w: %v", ErrFatal, reason)
	}

	return
}
