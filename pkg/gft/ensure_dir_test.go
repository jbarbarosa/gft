package gft_test

import (
	"errors"
	"os"
	"testing"

	"github.com/jbarbarosa/gft/pkg/gft"
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

func TestShouldSwitchDirectoriesIfPathExists(t *testing.T) {
	currdir, _ := os.Getwd()
	path := "somewhere/over/the/"
	filename := "rainbow_test.go"
	if err := os.MkdirAll(path, 0777); err != nil {
		t.Fatalf("test setup: failed to create directories with path: %s, %s", path, err)
	}
	file, err := os.Create(path + filename)
	defer func() {
		os.Chdir(currdir)
		if err := file.Close(); err != nil {
			t.Fatalf("failed to cleanup file pointer: %s", err)
		}
		if err := os.Remove(path + filename); err != nil {
			t.Fatalf("failed to cleanup file: %s", err)
		}
		if err := os.RemoveAll("somewhere"); err != nil {
			t.Fatalf("failed to cleanup directories: %s", err)
		}
	}()

	if err != nil {
		t.Fatalf("test setup: failed to create file with path: %s, %s", path, err)
	}

	if err = gft.EnsureDirFromPath(path); err != nil {
		t.Fatalf("expected to switch to existing directory, but failed: %s", err)
	}

	if err := os.RemoveAll("somewhere"); err != nil {
		t.Fatalf("failed to cleanup directories: %s", err)
	}
}
