package client

import (
	"fmt"

	"github.com/plaid/plaid-go/plaid"
	"github.com/syllabix/rollpay/backend/config"
	"github.com/syllabix/rollpay/backend/util/retryable"
)

type Client struct {
	*plaid.Client
}

func optsFrom(settings config.PlaidSettings) plaid.ClientOptions {
	return plaid.ClientOptions{
		ClientID:    settings.ClientID,
		Secret:      settings.Secret,
		Environment: settings.Environment,
		HTTPClient:  retryable.NewClient(),
	}
}

func New(settings config.PlaidSettings) (Client, error) {
	client, err := plaid.NewClient(optsFrom(settings))
	if err != nil {
		return Client{}, fmt.Errorf("failed to initialize Plaid client: %w", err)
	}

	return Client{client}, nil
}
