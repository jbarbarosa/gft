package gft_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/jbarbarosa/gft/internal/gft"
)

func TestShouldListTestNamesFromFiles(t *testing.T) {
	file := "file_test.go"
	expected := []string{"TestShouldListTestNamesFromFiles", "TestShouldReturnErrorIfFileIsNotFound"}

	if got, _ := gft.FromFile(file); !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected test name: %s, got %s", expected, got)
	}
}

func TestShouldReturnErrorIfFileIsNotFound(t *testing.T) {
	expected := gft.ErrFileNotFound
	file := "gone.go"

	if _, got := gft.FromFile(file); !errors.Is(got, expected) {
		t.Fatal("expected an error due to no file found, got no errors")
	}
}
