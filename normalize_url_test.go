package main

import (
	"testing"
)

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name     string
		inputURL string
		expected string
	}{
		{
			name:     "remove scheme",
			inputURL: "https://blog.boot.dev/path/",
			expected: "blog.boot.dev/path/",
		},
		{
			name:     "remove scheme 2",
			inputURL: "http://blog.boot.dev/path/",
			expected: "blog.boot.dev/path/",
		},
		{
			name:     "remove scheme and add trailing slash",
			inputURL: "https://blog.boot.dev/path",
			expected: "blog.boot.dev/path/",
		},
		{
			name:     "remove scheme and add trailing slash 2",
			inputURL: "http://blog.boot.dev/path",
			expected: "blog.boot.dev/path/",
		},
		{
			name:     "already parsed",
			inputURL: "blog.boot.dev/path/",
			expected: "blog.boot.dev/path/",
		},
		{
			name:     "add trailing slash",
			inputURL: "blog.boot.dev/path",
			expected: "blog.boot.dev/path/",
		},
		{
			name:     "no path",
			inputURL: "https://blog.boot.dev/",
			expected: "blog.boot.dev/",
		},
		{
			name:     "no path 2",
			inputURL: "https://blog.boot.dev",
			expected: "blog.boot.dev/",
		},
		// add more test cases here
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := normalizeURL(tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
