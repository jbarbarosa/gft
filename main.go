package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/jbarbarosa/gft/internal/gft"
)

func main() {
	path, file := gft.SplitPathFile(os.Args[1])
	if len(file) < 1 {
		log.Fatalf("fatal: go test file not found in path: %s", path)
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
