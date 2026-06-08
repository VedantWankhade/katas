package slidingwindow_test

import (
	"testing"

	"github.com/vedantwankhade/katas/leetcode/top-interview-150/slidingwindow"
)

func TestLongestSubstringWithoutRepeatingCharacters(t *testing.T) {
	tests := []struct {
		s        string
		expected int
	}{
		{
			s:        "abcabcbb",
			expected: 3,
		},
		{
			s:        "bbbbb",
			expected: 1,
		},
		{
			s:        "pwwkew",
			expected: 3,
		},
		{
			s:        "au",
			expected: 2,
		},
		{
			s:        "abba",
			expected: 2,
		},
		{
			s:        "tmmzuxt",
			expected: 5,
		},
	}

	for _, test := range tests {
		if actual := slidingwindow.LongestSubstringWithoutRepeatingCharacters(test.s); actual != test.expected {
			t.Errorf("\nS: %v\nExpected: %v\tActual: %v\n", test.s, test.expected, actual)
		}
	}
}
