package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/jbarbarosa/gft/internal/gft"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("go file test: pass the name of a go test file to run all its tests")
		os.Exit(1)
	}
	switch args := os.Args[1]; args {
	case "version":
		fmt.Println("gft version 0.1.0")

	default:
		path, file := gft.SplitPathFile(args)
		if len(file) < 1 {
			fmt.Println("fatal: go test file not found in path:", path)
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
