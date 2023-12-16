package main

import (
	"fmt"
	"os"

	"github.com/jbarbarosa/gft/cmd"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("go file test: pass the name of a go test file to run all its tests")
		os.Exit(1)
	}
	switch os.Args[1] {
	case "version":
		fmt.Println("gft version 0.1.0")
	default:
		cmd.Run(os.Args[1])
	}
}
