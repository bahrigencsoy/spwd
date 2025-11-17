package main

import (
	"os"
	"testing"
)

func TestShortenPath(t *testing.T) {
	// Mock home directory for testing
	originalHome := os.Getenv("HOME")
	os.Setenv("HOME", "/home/john")
	defer os.Setenv("HOME", originalHome)

	// Force UserHomeDir to return our test home
	testCases := []struct {
		input    string
		expected string
	}{
		// Basic cases
		{"/opt/xyz", "/opt/xyz"},
		{"/opt/xyz/klm", "/opt/xyz/klm"},
		{"/opt/xyz/klm/abc", "/opt/.../klm/abc"},
		{"/opt/xyz/klm/abc/123", "/opt/..../abc/123"},
		{"/tmp", "/tmp"},

		// Home directory cases
		{"/home/john", "~"},
		{"/home/john/projects", "~/projects"},
		{"/home/john/projects/abc", "~/projects/abc"},
		{"/home/john/projects/abc/target", "~/.../abc/target"},

		// More complex cases
		{"/opt/xyz/klm/abc/123/456", "/opt/...../123/456"},
		{"/home/john/a/b/c/d", "~/.../c/d"},

		// Edge cases
		{"/", "/"},
		{"/usr", "/usr"},
		{"/usr/local", "/usr/local"},
		{"/usr/local/bin", "/usr/local/bin"},
		{"/usr/local/bin/test", "/usr/.../bin/test"},

		// Home with more levels
		{"/home/john/a", "~/a"},
		{"/home/john/a x y/b", "~/a x y/b"},
		{"/home/john/a/b/c", "~/.../b/c"},
		{"/home/john/a/b/c/d/e f g", "~/..../d/e f g"},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := ShortenPath(tc.input)
			if result != tc.expected {
				t.Errorf("ShortenPath(%q) = %q; want %q", tc.input, result, tc.expected)
			}
		})
	}
}

func TestSplitPath(t *testing.T) {
	testCases := []struct {
		input    string
		expected []string
	}{
		{"/", []string{"/"}},
		{"/opt", []string{"/", "opt"}},
		{"/opt/xyz", []string{"/", "opt", "xyz"}},
		{"~", []string{"~"}},
		{"~/projects", []string{"~", "projects"}},
		{"~/projects/abc", []string{"~", "projects", "abc"}},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := splitPath(tc.input)
			if len(result) != len(tc.expected) {
				t.Errorf("splitPath(%q) length = %d; want %d", tc.input, len(result), len(tc.expected))
				return
			}
			for i := range result {
				if result[i] != tc.expected[i] {
					t.Errorf("splitPath(%q)[%d] = %q; want %q", tc.input, i, result[i], tc.expected[i])
				}
			}
		})
	}
}

func TestJoinPath(t *testing.T) {
	testCases := []struct {
		input    []string
		expected string
	}{
		{[]string{"/"}, "/"},
		{[]string{"/", "opt"}, "/opt"},
		{[]string{"/", "opt", "xyz"}, "/opt/xyz"},
		{[]string{"~"}, "~"},
		{[]string{"~", "projects"}, "~/projects"},
		{[]string{"~", "projects", "abc"}, "~/projects/abc"},
		{[]string{"/", "...", "bin", "test"}, "/.../.../test"},
	}

	for _, tc := range testCases {
		t.Run(tc.expected, func(t *testing.T) {
			result := joinPath(tc.input)
			if result != tc.expected {
				t.Errorf("joinPath(%v) = %q; want %q", tc.input, result, tc.expected)
			}
		})
	}
}

// Benchmark tests
func BenchmarkShortenPath(b *testing.B) {
	os.Setenv("HOME", "/home/john")
	testPath := "/home/john/projects/abc/target/debug/build/something/else"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ShortenPath(testPath)
	}
}

func BenchmarkShortenPathShort(b *testing.B) {
	os.Setenv("HOME", "/home/john")
	testPath := "/opt/xyz"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ShortenPath(testPath)
	}
}
