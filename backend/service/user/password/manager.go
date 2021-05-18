package password

import (
	"errors"
	"fmt"

	"github.com/syllabix/rollpay/backend/config"
	"golang.org/x/crypto/bcrypt"
)

var (
	// ErrSizeExceeded is returned if the the provided password is too large
	// If this error is returned - a malicious attempt to DOS the account setup process should
	// be suspected
	ErrSizeExceeded = errors.New("the provided password exceededs the maximum allowable size")
)

// Manager is used for hashing and comparing
// passwords
type Manager struct {
	cost      int
	maxLength int
}

// GenerateHash will create a secure hash of the provided password
func (mngr Manager) GenerateHash(password string) ([]byte, error) {
	if len(password) > mngr.maxLength {
		return []byte{}, ErrSizeExceeded
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), mngr.cost)
	if err != nil {
		return []byte{}, fmt.Errorf("failed to create password: %w", err)
	}

	return hash, nil
}

// Compare compares a hashed password with its possible plaintext equivalent.
// Returns nil on success, or an error on failure.
func (mngr Manager) Compare(hash []byte, password string) error {
	if len(password) > mngr.maxLength {
		return ErrSizeExceeded
	}

	return bcrypt.CompareHashAndPassword(hash, []byte(password))
}

// NewManager constructs a password manager with the provided
// settings
func NewManager(settings config.SecuritySettings) Manager {
	return Manager{
		cost:      settings.PasswordHashCost,
		maxLength: settings.PasswordMaxLength,
	}
}
