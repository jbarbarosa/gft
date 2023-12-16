package gft

import (
	"errors"
	"fmt"
	"os"
)

var ErrFileNotFound = errors.New("no such file")

func OpenFile(name string, fn func(*os.File)) error {
	file, err := os.Open(name)
	defer file.Close()

	if err != nil {
		return fmt.Errorf("%s: %w", name, ErrFileNotFound)
	}

	fn(file)
	return nil
}
