package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestShortenPath(t *testing.T) {
	const homeDir = "/home/john"

	testCases := []struct {
		name     string
		path     string
		home     string
		expected string
	}{
		{
			name:     "Home directory",
			path:     "/home/john",
			home:     homeDir,
			expected: "~",
		},
		{
			name:     "Two-level absolute path",
			path:     "/opt/xyz",
			home:     homeDir,
			expected: "/opt/xyz",
		},
		{
			name:     "Three-level absolute path",
			path:     "/opt/xyz/klm",
			home:     homeDir,
			expected: "/opt/xyz/klm",
		},
		{
			name:     "Four-level absolute path",
			path:     "/opt/xyz/klm/abc",
			home:     homeDir,
			expected: "/opt/.../klm/abc",
		},
		{
			name:     "Five-level absolute path",
			path:     "/opt/xyz/klm/abc/123",
			home:     homeDir,
			expected: "/opt/..../abc/123",
		},
		{
			name:     "Tmp directory",
			path:     "/tmp",
			home:     homeDir,
			expected: "/tmp",
		},
		{
			name:     "One level inside home",
			path:     "/home/john/projects",
			home:     homeDir,
			expected: "~/projects",
		},
		{
			name:     "Two levels inside home",
			path:     "/home/john/projects/abc",
			home:     homeDir,
			expected: "~/projects/abc",
		},
		{
			name:     "Three levels inside home",
			path:     "/home/john/projects/abc/target",
			home:     homeDir,
			expected: "~/.../abc/target",
		},
		{
			name:     "Root directory",
			path:     "/",
			home:     homeDir,
			expected: "/",
		},
		{
			name:     "No home dir match",
			path:     "/home/jane/projects",
			home:     homeDir,
			expected: "/home/jane/projects",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := ShortenPath(tc.path, tc.home)
			if got != tc.expected {
				t.Errorf("ShortenPath(%q, %q) = %q; want %q", tc.path, tc.home, got, tc.expected)
			}
		})
	}
}

func TestTFlag(t *testing.T) {
	ppid := os.Getpid()
	tmpDir := os.TempDir()
	filePath := filepath.Join(tmpDir, fmt.Sprintf("shortpath-%d", ppid))

	cmd := exec.Command("./shortpath", "-t")
	cmd.Run()

	// Check if the file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		t.Fatalf("file was not created: %s", filePath)
	}

	// Check file permissions
	info, err := os.Stat(filePath)
	if err != nil {
		t.Fatalf("failed to get file info: %v", err)
	}
	if info.Mode().Perm() != 0600 {
		t.Errorf("file permissions are not 0600: %s", info.Mode().Perm())
	}

	// Check file content
	content, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("failed to read file: %v", err)
	}
	if len(strings.TrimSpace(string(content))) == 0 {
		t.Errorf("file is empty")
	}

	// Clean up the file
	os.Remove(filePath)
}