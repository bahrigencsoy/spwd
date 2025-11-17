package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	useColor := flag.Bool("c", false, "use color in output")
	flag.Parse()

	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("failed to get current working directory: %v", err)
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		// Don't fail, just proceed without home dir shortening
		homeDir = ""
	}

	shortenedPath := ShortenPath(wd, homeDir)

	if *useColor {
		fmt.Printf("\033[31m%s\033[0m\n", shortenedPath)
	} else {
		fmt.Println(shortenedPath)
	}
}
