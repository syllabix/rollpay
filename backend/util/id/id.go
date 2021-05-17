package id

import (
	"errors"
	"fmt"
	"strconv"
)

var ErrInvalid = errors.New("the provided id is not a valid internal id")

// As public converts an internal int64 id
// to a string
func AsPublic(id int64) string {
	return strconv.FormatInt(id, 10)
}

// ToInternal converts the provided id into a valid internal numeric
// id, returning a non nil error if the id is invalid
func ToInternal(id string) (int64, error) {
	internalID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("%w: %v", ErrInvalid, err)
	}
	return internalID, nil
}
