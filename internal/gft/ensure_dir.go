package gft

import "os"

func EnsureDirFromPath(path string) error {
	return os.Chdir(path)
}
