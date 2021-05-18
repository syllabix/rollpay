package retryable

import "net/http"

type transport struct {
	roundtipper http.RoundTripper
}

func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("User-Agent", "rollpay/v0.0.1")
	return t.roundtipper.RoundTrip(req)
}

func newTransport() *transport {
	return &transport{
		roundtipper: http.DefaultTransport,
	}
}
