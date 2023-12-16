package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/jbarbarosa/gft/pkg/gft"
)

func Run(arg string) {
	path, file := gft.SplitPathFile(arg)
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
