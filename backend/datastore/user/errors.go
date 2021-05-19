package user

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/lib/pq"
	"github.com/syllabix/rollpay/backend/datastore/model"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

var (
	ErrNotFound   = errors.New("the requested user was not found")
	ErrFatal      = errors.New("an unexpected error occurred while accessing the database")
	ErrEmailTaken = errors.New("the email being used to create a new account is already in use")
)

func mapErr(reason error) (err error) {
	var sqlErr *pq.Error
	switch {
	case errors.Is(reason, sql.ErrNoRows):
		err = ErrNotFound

	case errors.As(reason, &sqlErr):
		if sqlErr.Code.Name() == "unique_violation" {
			err = ErrEmailTaken
			return
		}
		// if it is not an expected error, then use defaul case
		fallthrough

	default:
		err = fmt.Errorf("%w: %v", ErrFatal, reason)
	}

	return
}

func failure(reason error) (user model.User, err error) {
	return user, mapErr(reason)
}

func rollback(tx boil.Transactor, reason error) (user model.User, err error) {
	err = tx.Rollback()
	if err != nil {
		return user, mapErr(err)
	}

	return user, mapErr(reason)
}
