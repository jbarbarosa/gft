package gft_test

import (
	"reflect"
	"testing"

	"github.com/jbarbarosa/gft/internal/gft"
)

func TestShouldSplitPathFromFile(t *testing.T) {
	for _, test := range []struct {
		name     string
		fullpath string
		expected [2]string
	}{
		{
			name:     "should return path and empty file when only dirs are passed",
			fullpath: "somewhere/over/the/rainbow",
			expected: [2]string{"somewhere/over/the/rainbow", ""},
		},
		{
			name:     "should return path and empty file when the it does not match *_test.go",
			fullpath: "somewhere/over/the/rainbow.txt",
			expected: [2]string{"somewhere/over/the/rainbow.txt", ""},
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			if got := gft.SplitPathFile(test.fullpath); !reflect.DeepEqual(got, test.expected) {
				t.Fatalf("slice does not match expectation: got %v, want %v", got, test.expected)
			}
		})
	}
}
