package main

import (
	"strings"
)

// ShortenPath shortens a file path according to specific rules.
func ShortenPath(path string, homeDir string) string {
	if path == "/" {
		return "/"
	}

	var segments []string
	isHome := false

	if homeDir != "" && strings.HasPrefix(path, homeDir) {
		if path == homeDir {
			return "~"
		}
		if strings.HasPrefix(path, homeDir+"/") {
			isHome = true
			// path will be like "projects/shortpath"
			path = strings.TrimPrefix(path, homeDir+"/")
			segments = append([]string{"~"}, strings.Split(path, "/")...)
		}
	}

	if !isHome {
		path = strings.TrimPrefix(path, "/")
		segments = strings.Split(path, "/")
	}

	// For path "/a/b", segments is ["a", "b"] (len 2).
	// For path "~/a/b", segments is ["~", "a", "b"] (len 3).
	// The rule is to show at most 3 elements: root, and last two.
	if len(segments) <= 3 {
		if isHome {
			return strings.Join(segments, "/")
		}
		return "/" + strings.Join(segments, "/")
	}

	// More than 3 segments, shortening is needed.
	// ["a", "b", "c", "d"] -> a/.../c/d
	// ["~", "b", "c", "d"] -> ~/.../c/d
	middlePartsCount := len(segments) - 3 // 1 for root, 2 for last two parts
	dots := strings.Repeat(".", middlePartsCount+2)

	resultSegments := []string{
		segments[0],
		dots,
		segments[len(segments)-2],
		segments[len(segments)-1],
	}

	if isHome {
		return strings.Join(resultSegments, "/")
	}

	return "/" + strings.Join(resultSegments, "/")
}