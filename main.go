package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	findGit := flag.Bool("g", false, "traverse up and find the closest .git directory")
	flag.Parse()

	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("failed to get current working directory: %v", err)
	}

	if *findGit {
		repoName := FindClosestGitRepoParent(wd)
		if repoName != "" {
			fmt.Printf("[%s] ", repoName)
		}
		return
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		// Don't fail, just proceed without home dir shortening
		homeDir = ""
	}

	shortenedPath := ShortenPath(wd, homeDir)

	fmt.Println(shortenedPath)
}
