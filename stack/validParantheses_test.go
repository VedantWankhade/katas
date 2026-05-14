package stack_test

import (
	"testing"

	"github.com/vedantwankhade/katas/leetcode/top-interview-150/stack"
)

func TestValidParantheses(t *testing.T) {
	tests := []struct {
		s        string
		expected bool
	}{
		{
			s:        "()",
			expected: true,
		},
		{
			s:        "()[]{}",
			expected: true,
		},
		{
			s:        "(]",
			expected: false,
		},
		{
			s:        "([])",
			expected: true,
		},
		{
			s:        "([)]",
			expected: false,
		},
		{
			s:        "(){}}{",
			expected: false,
		},
	}

	for _, test := range tests {
		if actual := stack.ValidParantheses(test.s); actual != test.expected {
			t.Errorf("\nS: %v\nExpected: %v\tActual: %v\n", test.s, test.expected, actual)
		}
	}
}
