package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/jbarbarosa/gft/internal/gft"
)

func main() {
	gft.OpenFile(os.Args[1], func(file *os.File) {
		tests := gft.TestsFromFile(file)
		regex := gft.CreateRegex(tests)
		exec := exec.Command("go", "test", "-run", regex, ".")
		out, err := exec.Output()
		if err != nil {
			fmt.Println("error running tests: ", err)
		}
		fmt.Println(string(out))
	})
}
