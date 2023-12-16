package gft_test

import (
	"testing"

	"github.com/jbarbarosa/gft/internal/gft"
)

func TestShouldReturnAnErrorIfDirectoryDoesNotExist(t *testing.T) {
	err := gft.EnsureDirFromPath("idont/really/exist.go")
	if err == nil {
		t.Fatalf("expected an error, got no error")
	}
}
