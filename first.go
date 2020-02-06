package helptest

import (
	"io/ioutil"
	"os"
	"testing"
)

// TempDir creates a new temporary directory with the given
// prefix and returns that directory's path and a function
// to delete it when the test is finished.
func TempDir(t *testing.T, prefix string) (string, func()) {
	t.Helper()
	directory, err := ioutil.TempDir("", prefix)
	if err != nil {
		t.Fatal(err)
	}
	cleanup := func() {
		os.RemoveAll(directory)
	}
	return directory, cleanup
}
