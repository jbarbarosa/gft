package gft_test

import (
	"errors"
	"os"
	"testing"

	"github.com/jbarbarosa/gft/internal/gft"
)

func TestShouldReturnOpenFileIfItExists(t *testing.T) {
	os.Create("testfile.go")
	defer func() { os.Remove("testfile.go") }()
	counter := 0
	err := gft.OpenFile("testfile.go", func(*os.File) {
		counter++
	})
	if err != nil {
		t.Fatalf("expected no errors, got %s", err)
	}

	if counter != 1 {
		t.Fatalf("expected number of calls to handler function to be 1, got %d", counter)
	}
}

func TestShouldReturnAnErrorIfTheFileWasNotFound(t *testing.T) {
	counter := 0
	err := gft.OpenFile("testfile.go", func(*os.File) {
		counter++
	})

	if !errors.Is(err, gft.ErrFileNotFound) {
		t.Fatalf("expected error of type errfilenotfound, got %s", err)
	}
}
