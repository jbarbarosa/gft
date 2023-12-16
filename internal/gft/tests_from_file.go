package gft

import (
	"bufio"
	"os"
	"strings"
)

// TestsFromFile returns all tests in the file passed to it
func TestsFromFile(file *os.File) []string {
	testNames := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if i := strings.Index(line, "func Test"); i != -1 {
			startIndex := i + 5
			parenthesesIndex := strings.Index(line, "(")
			testNames = append(testNames, line[startIndex:parenthesesIndex])
		}
	}
	return testNames
}
