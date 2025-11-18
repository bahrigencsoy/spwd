package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {

	findGit := flag.Bool("g", false, "traverse up and find the closest .git directory")
	micro := flag.Bool("m", false, "current microsecond clock")
	diff := flag.Int64("d", 0, "microsecond diff")
	flag.Parse()

	if *micro {
		fmt.Println(time.Now().UnixMicro())
		return
	}

	if *diff != 0 {
		resultFloat := float64(time.Now().UnixMicro()-*diff) / 1000000.0
		fmt.Printf("%.3f\n", resultFloat)
		return
	}

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
