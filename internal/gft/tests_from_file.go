package gft

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

var ErrFileNotFound = errors.New("no such test file or directory")

// TestsFromFile returns all tests in the file passed to it
func TestsFromFile(filename string) ([]string, error) {
	testNames := []string{}
	file, err := os.Open(filename)
	if err != nil {
		return testNames, ErrFileNotFound
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if i := strings.Index(line, "func Test"); i != -1 {
			startIndex := i + 5
			parenthesesIndex := strings.Index(line, "(")
			testNames = append(testNames, line[startIndex:parenthesesIndex])
		}
	}
	return testNames, nil
}
