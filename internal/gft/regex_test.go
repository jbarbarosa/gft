package gft_test

import (
	"testing"

	"github.com/jbarbarosa/gft/internal/gft"
)

func TestShouldCreateRegexFromTestNames(t *testing.T) {
	names := []string{"TestShouldListTestNamesFromFiles", "TestShouldReturnErrorIfFileIsNotFound"}
	expected := "^Test(ShouldListTestNamesFromFiles|ShouldReturnErrorIfFileIsNotFound)$"

	if got := gft.Regex(names); got != expected {
		t.Fatalf("expected regex to match: %s, got %s", expected, got)
	}
}
