package gft

import "strings"

func SplitPathFile(path string) [2]string {
	lastslash := strings.LastIndex(path, "/")
	if strings.Index(path[lastslash:], "_test.go") == -1 {
		return [2]string{path, ""}
	}
	return [2]string{path[:lastslash+1], path[lastslash+1:]}
}
