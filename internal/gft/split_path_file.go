package gft

import "strings"

func SplitPathFile(fullpath string) (path string, file string) {
	lastslash := strings.LastIndex(fullpath, "/")
	if strings.Index(fullpath[lastslash:], "_test.go") == -1 {
		path = fullpath
		return
	}
	path = fullpath[:lastslash+1]
	file = fullpath[lastslash+1:]
	return
}
