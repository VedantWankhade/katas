package arraystring_test

import (
	"testing"

	"github.com/vedantwankhade/katas/leetcode/top-interview-150/arraystring"
)

func TestLengthOfLastWordUsingBuiltInMethods(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			input:    "Hello World",
			expected: 5,
		},
		{
			input:    "   fly me   to   the moon  ",
			expected: 4,
		},
		{
			input:    "luffy is still joyboy",
			expected: 6,
		},
	}

	for _, test := range tests {
		actual := arraystring.LengthOfLastWordUsingBuiltInMethods(test.input)
		if actual != test.expected {
			t.Errorf("\nWord: %v\nExpected: %v\tActual: %v", test.input, test.expected, actual)
		}
	}
}

func TestLengthOfLastWord(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			input:    "Hello World",
			expected: 5,
		},
		{
			input:    "   fly me   to   the moon  ",
			expected: 4,
		},
		{
			input:    "luffy is still joyboy",
			expected: 6,
		},
	}

	for _, test := range tests {
		actual := arraystring.LengthOfLastWord(test.input)
		if actual != test.expected {
			t.Errorf("\nWord: %v\nExpected: %v\tActual: %v", test.input, test.expected, actual)
		}
	}
}
