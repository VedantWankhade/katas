package arraystring_test

import (
	"testing"

	"github.com/vedantwankhade/katas/leetcode/top-interview-150/arraystring"
)

func TestRomanToInteger(t *testing.T) {
	tests := []struct {
		roman    string
		expected int
	}{
		{
			roman:    "III",
			expected: 3,
		},
		{
			roman:    "LVIII",
			expected: 58,
		},
		{
			roman:    "MCMXCIV",
			expected: 1994,
		},
	}
	for _, test := range tests {
		actual := arraystring.RomanToInteger(test.roman)
		if actual != test.expected {
			t.Errorf("Roman: %v\nExpected: %v\tActual: %v", test.roman, test.expected, actual)
		}
	}
}
