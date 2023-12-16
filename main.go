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
	regex := gft.CreateRegex(tests)
	exec := exec.Command("go", "test", "-run", regex, ".")
	out, err := exec.Output()
	if err != nil {
		fmt.Println("error running tests: ", err)
	}
	fmt.Println(string(out))
}
