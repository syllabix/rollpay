package retryable

import (
	"net/http"
	"time"

	"github.com/hashicorp/go-retryablehttp"
)

// retry defaults
const (
	id      = "rollpay"
	timeout = time.Second * 5
	retries = 5
)

type config struct {
	timeout time.Duration
	retries int
}

// An Option is used to configure an instance of a retryable http client
type Option func(*config)

// Timeout is an option that sets the retryable client timeout
func Timeout(duration time.Duration) Option {
	return func(c *config) {
		c.timeout = duration
	}
}

// MaxRetries is an option that sets the retryable client max
// retry count
func MaxRetries(count int) Option {
	return func(c *config) {
		c.retries = count
	}
}

// Client intializes a standard http.Client with retries, backoff and
// timeouts using the provided options
func NewClient(options ...Option) *http.Client {
	c := &config{timeout, retries}

	for _, opt := range options {
		opt(c)
	}

	retryClient := retryablehttp.NewClient()
	retryClient.Logger = nil
	retryClient.RetryMax = c.retries
	retryClient.HTTPClient.Timeout = c.timeout
	retryClient.HTTPClient.Transport = newTransport()

	return retryClient.StandardClient()
}
