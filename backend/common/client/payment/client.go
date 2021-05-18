package payment

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/plaid/plaid-go/plaid"
	"github.com/syllabix/rollpay/backend/common/retryable"
	"github.com/syllabix/rollpay/backend/config"
)

// A User that initiaties various client payment requests
type User struct {
	ID       string
	Email    string
	Language string
}

type LinkToken struct {
	Token      string
	RequestID  string
	Expiration time.Time
}

// Client is used to interface with our third party payment
// provider
type Client struct {
	client *plaid.Client

	products     []string
	redirectURI  string
	countryCodes []string
	clientName   string
}

// Create a link token for the provided payment user
func (c *Client) CreateLinkToken(ctx context.Context, user User) (LinkToken, error) {
	configs := c.tokenConfigsFor(user)
	resp, err := c.client.CreateLinkToken(configs)
	if err != nil {
		return LinkToken{}, fmt.Errorf("failed to create link token: %w", err)
	}

	return LinkToken{
		Token:      resp.LinkToken,
		RequestID:  resp.RequestID,
		Expiration: resp.Expiration,
	}, nil
}

func (c *Client) tokenConfigsFor(user User) plaid.LinkTokenConfigs {
	return plaid.LinkTokenConfigs{
		User: &plaid.LinkTokenUser{
			ClientUserID: user.ID,
		},
		ClientName:   c.clientName,
		Products:     c.products,
		CountryCodes: c.countryCodes,
		Language:     user.Language,
		RedirectUri:  c.redirectURI,
	}
}

func NewClient(settings config.PlaidSettings) (*Client, error) {
	opts := optsFrom(settings)
	client, err := plaid.NewClient(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Plaid client: %w", err)
	}

	return &Client{
		client:       client,
		products:     strings.Split(settings.Products, ","),
		redirectURI:  settings.RedirectURI,
		countryCodes: strings.Split(settings.CountryCodes, ","),
		clientName:   settings.ClientName,
	}, nil
}

func optsFrom(settings config.PlaidSettings) plaid.ClientOptions {
	return plaid.ClientOptions{
		ClientID:    settings.ClientID,
		Secret:      settings.Secret,
		Environment: settings.Environment,
		HTTPClient:  retryable.NewClient(),
	}
}
