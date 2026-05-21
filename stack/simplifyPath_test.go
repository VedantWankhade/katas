package stack_test

import (
	"testing"

	"github.com/vedantwankhade/katas/leetcode/top-interview-150/stack"
)

func TestSimplifyPath(t *testing.T) {
	tests := []struct {
		path     string
		expected string
	}{
		{
			path:     "/home/",
			expected: "/home",
		},
		{
			path:     "/home///foo/",
			expected: "/home/foo",
		},
		{
			path:     "/home/user/Documents/../Pictures",
			expected: "/home/user/Pictures",
		},
		{
			path:     "/../",
			expected: "/",
		},
		{
			path:     "/.../a/../b/c/../d/./",
			expected: "/.../b/d",
		},
	}

	for _, test := range tests {
		if actual := stack.SimplifyPath(test.path); actual != test.expected {
			t.Errorf("\nPath: %v\nExpected: %v\tActual: %v\n", test.path, test.expected, actual)
		}
	}
}
