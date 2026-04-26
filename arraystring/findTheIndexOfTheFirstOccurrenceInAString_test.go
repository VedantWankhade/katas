package arraystring_test

import (
	"testing"

	"github.com/vedantwankhade/katas/leetcode/top-interview-150/arraystring"
)

func TestFindTheIndexOfTheFirstOccurrenceInAString(t *testing.T) {
	tests := []struct {
		haystack string
		needle   string
		expected int
	}{
		{
			haystack: "sadbutsad",
			needle:   "sad",
			expected: 0,
		},
		{
			haystack: "leetcode",
			needle:   "leeto",
			expected: -1,
		},
		{
			haystack: "mississippi",
			needle:   "issip",
			expected: 4,
		},
	}

	for _, test := range tests {
		actual := arraystring.FindTheIndexOfTheFirstOccurrenceInAString(test.haystack, test.needle)
		if actual != test.expected {
			t.Errorf("\nHaystack: %v\tNeedle: %v\nExpected: %v\tActual: %v\n", test.haystack, test.needle, test.expected, actual)
		}
	}
}
