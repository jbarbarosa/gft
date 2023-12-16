package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/jbarbarosa/gft/internal/gft"
)

func main() {
	switch arg := os.Args[1]; arg {

	case "version":
		fmt.Println("gft version 0.1.0")

	default:
		path, file := gft.SplitPathFile(os.Args[1])
		if len(file) < 1 {
			fmt.Println("fatal: go test file not found in path: %s", path)
			os.Exit(1)
		}
		gft.EnsureDirFromPath(path)
		gft.OpenFile(file, func(file *os.File) {
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
}
