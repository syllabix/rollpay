package organization

import (
	"errors"
	"fmt"

	"github.com/syllabix/rollpay/backend/api/model"
	"github.com/syllabix/rollpay/backend/common/id"
	"github.com/syllabix/rollpay/backend/datastore/organization"
)

var (
	ErrFatal        = errors.New("an unexpected system error occurred")
	ErrNotFound     = errors.New("the requested organization does not exist")
	ErrInvalid      = errors.New("the request is invalid")
	ErrNameReserved = errors.New("the organization name is already in use")
)

func mapErr(reason error) (err error) {
	switch {
	case errors.Is(reason, organization.ErrNotFound):
		err = ErrNotFound

	case errors.Is(reason, organization.ErrNameTaken):
		err = ErrNameReserved

	case errors.Is(reason, id.ErrInvalid):
		err = fmt.Errorf("%w: %v", ErrInvalid, reason)

	default:
		err = fmt.Errorf("%v: %w", ErrFatal, reason)
	}

	return err
}

func failure(reason error) (u model.Organization, err error) {
	err = mapErr(reason)
	return
}
