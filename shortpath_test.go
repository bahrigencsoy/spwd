package main

import (
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