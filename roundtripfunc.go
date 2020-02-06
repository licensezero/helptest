package helptest

import (
	"net/http"
)

// RoundTripFunc implements RoundTrip with a single function.
type RoundTripFunc func(req *http.Request) *http.Response

// RoundTrip calls the provided function.
func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}
