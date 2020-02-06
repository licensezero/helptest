package helptest

import (
	"os"
	"testing"
)

// Chdir changes the working directory to the provided path
// and returns a function that resets the working directory
// to its previous value.
func Chdir(t *testing.T, directory string) func() {
	t.Helper()
	prior, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err = os.Chdir(directory); err != nil {
		t.Fatal(err)
	}
	return func() {
		os.Chdir(prior)
	}
}
