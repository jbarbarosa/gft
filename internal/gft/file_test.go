package gft_test

import (
	"testing"

	"github.com/jbarbarosa/gft/internal/gft"
)

func TestShouldListTestNamesFromFiles(t *testing.T) {
	file := "file_test"
	expected := "TestShouldListTestNamesFromFiles"

	if got := gft.FromFile(file); got != expected {
		t.Fatalf("expected test name: %s, got %s", expected, got)
	}
}
