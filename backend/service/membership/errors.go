package membership

import (
	"errors"
	"fmt"

	"github.com/syllabix/rollpay/backend/api/model"
	"github.com/syllabix/rollpay/backend/datastore/membership"
)

var (
	ErrFatal    = errors.New("an unexpected system error occurred")
	ErrInvalid  = errors.New("the request to create a membership is invalid")
	ErrNotFound = errors.New("the requested membership does not exist")
)

func mapErr(reason error) (err error) {
	switch {
	case errors.Is(reason, membership.ErrNotFound):
		err = ErrNotFound

	case errors.Is(reason, membership.ErrInvalidMembership):
		err = fmt.Errorf("%v: %w", ErrInvalid, reason)

	default:
		err = fmt.Errorf("%v: %w", ErrFatal, reason)
	}

	return err
}

func failure(reason error) (u model.OrganizationMember, err error) {
	err = mapErr(reason)
	return
}
