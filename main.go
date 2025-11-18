package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

func main() {
	findGit := flag.Bool("g", false, "traverse up and find the closest .git directory")
	printParentPid := flag.Bool("p", false, "print parent process pid")
	trackTime := flag.Bool("t", false, "track time in a file in the temp directory")
	flag.Parse()

	if *trackTime {
		ppid := os.Getppid()
		tmpDir := os.TempDir()
		filePath := filepath.Join(tmpDir, fmt.Sprintf("shortpath-%d", ppid))

		f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
		if err != nil {
			log.Fatalf("failed to open file: %v", err)
		}
		defer f.Close()

		timestamp := time.Now().Format("2006-01-02 15:04:05.999999")
		if _, err := f.WriteString(timestamp + "\n"); err != nil {
			log.Fatalf("failed to write to file: %v", err)
		}
		return
	}

	if *printParentPid {
		fmt.Println(os.Getppid())
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
