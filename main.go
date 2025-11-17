package main

import (
	"fmt"
	"os"
)

func main() {
	// Get the current working directory (where the command was executed from)
	// This behaves like 'pwd -L' (logical path, not physical)
	pwd := os.Getenv("PWD")
	if pwd == "" {
		var err error
		pwd, err = os.Getwd()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting current directory: %v\n", err)
			os.Exit(1)
		}
	}

	shortened := ShortenPath(pwd)
	fmt.Println(shortened)
}
