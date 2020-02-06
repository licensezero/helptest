package helptest

import (
	"io"
)

// NoopCloser implements Closer without doing anything on
// close.  This is useful for creating Body for http.Response
// from a BufferString.
type NoopCloser struct {
	io.Reader
}

// Close does not thing and never fails.
func (NoopCloser) Close() error {
	return nil
}
