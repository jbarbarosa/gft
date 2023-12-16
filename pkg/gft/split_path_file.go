package gft

import "strings"

func SplitPathFile(fullpath string) (path string, file string) {
	path = fullpath
	lastslash := strings.LastIndex(fullpath, "/")
	if lastslash == -1 {
		return
	}
	if strings.Index(fullpath[lastslash:], "_test.go") == -1 {
		return
	}
	path = fullpath[:lastslash+1]
	file = fullpath[lastslash+1:]
	return
}
