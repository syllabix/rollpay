package plaidenv

import (
	"github.com/plaid/plaid-go/plaid"
)

var environments = map[string]plaid.Environment{
	"sandbox":     plaid.Sandbox,
	"development": plaid.Development,
	"production":  plaid.Production,
}

// Get returns plaid environment for the provided string.
// If the string is not a valid environment or is empty,
// the plaid Sandbox environment is returned
//
// Valid strings:
// sandbox, development, production
func Get(e string) plaid.Environment {
	env, exists := environments[e]
	if !exists {
		return plaid.Sandbox
	}
	return env
}
