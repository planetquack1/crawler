package main

import (
	"reflect"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{
		{
			name:     "no URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
		<html>
			<body>
			</body>
		</html>
		`,
			expected: []string{},
		},
		{
			name:     "absolute URLs, same as input",
			inputURL: "https://blog.boot.dev",
			inputBody: `
		<html>
			<body>
				<a href="https://blog.boot.dev/path/two">
					<span>Boot.dev</span>
				</a>
			</body>
		</html>
		`,
			expected: []string{"https://blog.boot.dev/path/two"},
		},
		{
			name:     "absolute URLs, different from input",
			inputURL: "https://blog.boot.dev",
			inputBody: `
		<html>
			<body>
				<a href="https://other.com/path/three">
					<span>Boot.dev</span>
				</a>
			</body>
		</html>
		`,
			expected: []string{"https://other.com/path/three"},
		},
		{
			name:     "relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
		<html>
			<body>
				<a href="/path/four">
					<span>Boot.dev</span>
				</a>
			</body>
		</html>
		`,
			expected: []string{"https://blog.boot.dev/path/four"},
		},
		{
			name:     "absolute and relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
		<html>
			<body>
				<a href="/path/one">
					<span>Boot.dev</span>
				</a>
				<a href="https://other.com/path/one">
					<span>Boot.dev</span>
				</a>
			</body>
		</html>
		`,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
		{
			name:     "all mix, absolute and relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
		<html>
			<body>
				<a href="/path/five">
					<span>Boot.dev</span>
				</a>
				<a href="https://other.com/path/five">
					<span>Boot.dev</span>
				</a>
				<a href="https://blog.boot.dev/path/five">
					<span>Boot.dev</span>
				</a>
			</body>
		</html>
		`,
			expected: []string{"https://blog.boot.dev/path/five", "https://other.com/path/five", "https://blog.boot.dev/path/five"},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getURLsFromHTML(tc.inputBody, tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
