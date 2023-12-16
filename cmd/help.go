package cmd

import "fmt"

func Help() {
	fmt.Println(`
    gft: wrapper around the standard go testing library

    To test a single file, run:
        gft file_test.go

    You can pass in a path to gft as well:
        gft internal/package/feature_test.go
  `)
}
