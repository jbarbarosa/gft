package gft_test

import (
	"os"
	"reflect"
	"testing"

	"github.com/jbarbarosa/gft/internal/gft"
)

func TestShouldListTestNamesFromFiles(t *testing.T) {
	expected := []string{"TestShouldListTestNamesFromFiles"}

	file, err := os.Open("tests_from_file_test.go")
	if err != nil {
		t.Fatalf("expected no errors, got %s", err)
	}
	names := gft.TestsFromFile(file)

	if !reflect.DeepEqual(names, expected) {
		t.Fatalf("expected test name: %s, got %s", expected, names)
	}
}
