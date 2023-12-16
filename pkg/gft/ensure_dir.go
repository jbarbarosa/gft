package gft

import (
	"errors"
	"fmt"
	"os"
)

var ErrDirNotFound = errors.New("provided directory not found")

func EnsureDirFromPath(path string) error {
	if err := os.Chdir(path); err != nil {
		return fmt.Errorf("ensure dir: %s, %w", path, ErrDirNotFound)
	}
	return nil
}
