package arraystring_test

import (
	"testing"

	"github.com/vedantwankhade/katas/leetcode/top-interview-150/arraystring"
)

func TestReverseWordsInAString(t *testing.T) {
	tests := []struct {
		s        string
		expected string
	}{
		{
			s:        "the sky is blue",
			expected: "blue is sky the",
		},
		{
			s:        "  hello world  ",
			expected: "world hello",
		},
		{
			s:        "a good   example",
			expected: "example good a",
		},
	}

	for _, test := range tests {
		actual := arraystring.ReverseWordsInAString(test.s)
		if actual != test.expected {
			t.Errorf("Input: %v\nExpected: %v\tActual: %v\n", test.s, test.expected, actual)
		}
	}
}
