# gft
Go file test: pass in a file to run all its tests at once, without tediously mapping dependencies

## Requirements
- Go toolchain

## Install
Simplest way is by running:
```
    go install github.com/jbarbarosa/gft@v0.1.1
```
Alternatively, clone the repository and run `make`

## Basic Usage
To use gft, simply pass in the path to the test file in which you want to run all tests
```
    gft path/to/file_test.go
```
gft is just a wrapper around the native testing library. It works by reading the file and forming a regex that matches exactly
the test names in that file; in effect running it without having to pass in manually all its dependencies

For questions related to the behavior of the testing library, refer to [its documentation](https://pkg.go.dev/testing)

Contributions are welcome
