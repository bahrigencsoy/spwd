package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
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
	fmt.Println(shortenedPath)
}