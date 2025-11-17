package main

import (
	"os"
	"path/filepath"
	"strings"
)

// ShortenPath shortens a directory path according to the specified rules
func ShortenPath(path string) string {
	if path == "" {
		return ""
	}

	// Normalize the path but keep it logical (don't resolve symlinks)
	path = filepath.Clean(path)

	// Get home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		homeDir = ""
	}

	// Replace home directory with ~
	originalPath := path
	if homeDir != "" && (path == homeDir || strings.HasPrefix(path, homeDir+string(filepath.Separator))) {
		if path == homeDir {
			return "~"
		}
		path = "~" + path[len(homeDir):]
	}

	// Split the path into components
	parts := splitPath(path)

	if len(parts) == 0 {
		return originalPath
	}

	// If path has 3 or fewer components, return as is
	if len(parts) <= 3 {
		return joinPath(parts)
	}

	// Path has more than 3 components
	// Keep: root, dots (representing middle), last 2
	root := parts[0]
	last2 := parts[len(parts)-2:]

	// Calculate number of hidden directories
	numHidden := len(parts) - 3 // total - (root + 2 last)

	// Create dots representation (minimum 3 dots)
	dots := strings.Repeat(".", numHidden+2)

	// Construct the shortened path
	result := []string{root, dots}
	result = append(result, last2...)

	return joinPath(result)
}

// splitPath splits a path into its components
func splitPath(path string) []string {
	if path == "" {
		return []string{}
	}

	// Handle root directory
	if path == "/" {
		return []string{"/"}
	}

	// Handle paths starting with ~
	if strings.HasPrefix(path, "~") {
		rest := strings.TrimPrefix(path, "~")
		rest = strings.TrimPrefix(rest, string(filepath.Separator))
		if rest == "" {
			return []string{"~"}
		}
		parts := []string{"~"}
		if rest != "" {
			parts = append(parts, strings.Split(rest, string(filepath.Separator))...)
		}
		return parts
	}

	// Handle absolute paths
	if strings.HasPrefix(path, "/") {
		rest := strings.TrimPrefix(path, "/")
		if rest == "" {
			return []string{"/"}
		}
		parts := []string{"/"}
		parts = append(parts, strings.Split(rest, string(filepath.Separator))...)
		return parts
	}

	// Handle relative paths (shouldn't happen in our use case, but handle anyway)
	return strings.Split(path, string(filepath.Separator))
}

// joinPath joins path components back into a path string
func joinPath(parts []string) string {
	if len(parts) == 0 {
		return ""
	}

	if len(parts) == 1 {
		return parts[0]
	}

	// Handle root directory cases
	if parts[0] == "/" {
		if len(parts) == 1 {
			return "/"
		}
		return "/" + strings.Join(parts[1:], "/")
	}

	if parts[0] == "~" {
		if len(parts) == 1 {
			return "~"
		}
		return "~/" + strings.Join(parts[1:], "/")
	}

	return strings.Join(parts, "/")
}
