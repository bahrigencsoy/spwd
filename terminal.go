package main

import (
	"fmt"
	"os"
	"time"

	"golang.org/x/term"
)

func DisplayHello() {
	// Get terminal dimensions
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		// Fallback if not a TTY or other error
		fmt.Println("HELLO")
		return
	}

	// ANSI escape codes
	const (
		saveCursor    = "\033[s"
		restoreCursor = "\033[u"
	)

	// Move to top right corner (line 1, column width - 5 for "HELLO")
	// The string "HELLO" has length 5.
	moveCursor := fmt.Sprintf("\033[1;%dH", width-5)

	// Print "HELLO"
	fmt.Print(saveCursor)
	fmt.Print(moveCursor)
	fmt.Print("HELLO")
	fmt.Print(restoreCursor)

	// Wait
	time.Sleep(2 * time.Second)

	// Clear the writing by overwriting with spaces
	fmt.Print(saveCursor)
	fmt.Print(moveCursor)
	fmt.Print("     ") // 5 spaces
	fmt.Print(restoreCursor)
}