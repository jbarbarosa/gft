package gft_test

import (
	"errors"
	"testing"

	"github.com/jbarbarosa/gft/internal/gft"
)

func TestShouldReturnAnErrorIfDirectoryDoesNotExist(t *testing.T) {
	err := gft.EnsureDirFromPath("idont/really/exist.go")
	if err == nil {
		t.Fatalf("expected an error, got no error")
	}
	if !errors.Is(err, gft.ErrDirNotFound) {
		t.Fatalf("expected error to be of type errdirnotfound, got: %s", err)
	}
}
