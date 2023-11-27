package gft

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// FromFile returns all test names in the file passed to it
func FromFile(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		return ""
	}

	scanner := bufio.NewScanner(file)
	var testName string
	for scanner.Scan() {
		line := scanner.Text()
		if i := strings.Index(line, "func Test"); i != -1 {
			startIndex := i + 5
			parenthesesIndex := strings.Index(line, "(")
			testName = line[startIndex:parenthesesIndex]
		}
	}
	return testName
}
