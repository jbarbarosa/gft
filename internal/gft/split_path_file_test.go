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
		expected []string
	}{
		{
			name:     "should return path and empty file when only dirs are passed",
			fullpath: "somewhere/over/the/rainbow",
			expected: []string{"somewhere/over/the/rainbow", ""},
		},
		{
			name:     "should return path and empty file when it does not match *_test.go",
			fullpath: "somewhere/over/the/rainbow.txt",
			expected: []string{"somewhere/over/the/rainbow.txt", ""},
		},
		{
			name:     "should return path and file when it matches *_test.go",
			fullpath: "somewhere/over/the/rainbow_test.go",
			expected: []string{"somewhere/over/the/", "rainbow_test.go"},
		},
		{
			name:     "should return path and empty string when it is not properly formatted as a path",
			fullpath: "somewhereovertherainbow_test.go",
			expected: []string{"somewhereovertherainbow_test.go", ""},
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			path, file := gft.SplitPathFile(test.fullpath)
			got := []string{path, file}
			if !reflect.DeepEqual(got, test.expected) {
				t.Fatalf("slice does not match expectation: got %v, want %v", got, test.expected)
			}
		})
	}
}
