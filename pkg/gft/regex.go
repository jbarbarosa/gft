package gft

import "strings"

func CreateRegex(testNames []string) string {
	regex := "^Test("
	for _, name := range testNames {
		name = strings.TrimPrefix(name, "Test")
		regex += name + "|"
	}
	regex = strings.TrimSuffix(regex, "|")
	return regex + ")$"
}
