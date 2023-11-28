package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/jbarbarosa/gft/internal/gft"
)

func main() {
	tests, err := gft.TestsFromFile(os.Args[1])
	if err != nil {
		fmt.Println("unable to run tests from file")
	}
	regex := gft.Regex(tests)
	fmt.Println("Running tests with the following regex:", regex)
	exec := exec.Command("go", "test", "-run", regex, ".")
	out, err := exec.Output()
	if err != nil {
		fmt.Println("error running tests: ", err)
	}
	fmt.Println(string(out))
}
