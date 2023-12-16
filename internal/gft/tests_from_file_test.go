package gft_test

import (
	"errors"
	"os"
	"reflect"
	"testing"

	"github.com/jbarbarosa/gft/internal/gft"
)

func TestShouldListTestNamesFromFiles(t *testing.T) {
	expected := []string{"TestShouldListTestNamesFromFiles", "TestShouldReturnErrorIfFileIsNotFound"}

	var names []string
	err := gft.OpenFile("tests_from_file_test.go", func(file *os.File) {
		names, _ = gft.TestsFromFile(file)
	})
	if err != nil {
		t.Fatalf("expected no errors, got %s", err)
	}

	if !reflect.DeepEqual(names, expected) {
		t.Fatalf("expected test name: %s, got %s", expected, names)
	}
}

func TestShouldReturnErrorIfFileIsNotFound(t *testing.T) {
	err := gft.OpenFile("gone.go", func(file *os.File) {
		// expected not to be called
		t.Fatal("expected an error due to no file found, got no errors")
	})

	if !errors.Is(err, gft.ErrFileNotFound) {
		t.Fatalf("expected error to be errfilenotfound, err is instead %s", err)
	}
}
