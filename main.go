package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	useColor := flag.Bool("c", false, "use color in output")
	findGit := flag.Bool("g", false, "traverse up and find the closest .git directory")
	doHello := flag.Bool("t", false, "display HELLO on the top right corner")
	flag.Parse()

	if *doHello {
		DisplayHello()
		return
	}

	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("failed to get current working directory: %v", err)
	}

	if *findGit {
		repoName := FindClosestGitRepoParent(wd)
		if repoName != "" {
			if *useColor {
				fmt.Printf("\033[33m[%s]\033[0m ", repoName)
			} else {
				fmt.Printf("[%s] ", repoName)
			}
		}
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
