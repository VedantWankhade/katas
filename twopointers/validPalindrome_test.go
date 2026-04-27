package twopointers_test

import (
	"testing"

	"github.com/vedantwankhade/katas/leetcode/top-interview-150/twopointers"
)

func TestValidPalindrome(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{
			input:    "A man, a plan, a canal: Panama",
			expected: true,
		},
		{
			input:    "race a car",
			expected: false,
		},
		{
			input:    " ",
			expected: true,
		},
		{
			input:    "0P",
			expected: false,
		},
		{
			input:    "ab_a",
			expected: true,
		},
	}

	for _, test := range tests {
		actual := twopointers.ValidPalindrome(test.input)
		if actual != test.expected {
			t.Errorf("\nInput: %v\nExpected: %v\tActual: %v\n", test.input, test.expected, actual)
		}
	}
}
