package main

import (
	"fmt"
	"os"

	"github.com/jbarbarosa/gft/cmd"
)

func main() {
	version := "0.1.0"
	if len(os.Args) < 2 || os.Args[1] == "help" || os.Args[1] == "--help" {
		cmd.Help()
		os.Exit(1)
	}
	switch os.Args[1] {
	case "version":
		fmt.Printf("gft version %s\n", version)
	default:
		cmd.Run(os.Args[1])
	}
}
