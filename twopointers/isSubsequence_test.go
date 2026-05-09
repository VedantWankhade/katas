package twopointers_test

import (
	"testing"

	"github.com/vedantwankhade/katas/leetcode/top-interview-150/twopointers"
)

func TestIsSubsequence(t *testing.T) {
	tests := []struct {
		s        string
		t        string
		expected bool
	}{
		{
			s:        "abc",
			t:        "ahbgdc",
			expected: true,
		},
		{
			s:        "axc",
			t:        "ahbgdc",
			expected: false,
		},
		{
			s:        "abc",
			t:        "ab",
			expected: false,
		},
	}
	for _, test := range tests {
		actual := twopointers.IsSubsequence(test.s, test.t)
		if actual != test.expected {
			t.Errorf("\nS: %v\tT: %v\nExpected: %v\tActual: %v", test.s, test.t, test.expected, actual)
		}
	}
}
