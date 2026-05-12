package hashmap_test

import (
	"testing"

	"github.com/vedantwankhade/katas/leetcode/top-interview-150/hashmap"
)

func TestValidAnagram(t *testing.T) {
	tests := []struct {
		s        string
		t        string
		expected bool
	}{
		{
			s:        "anagram",
			t:        "nagaram",
			expected: true,
		},
		{
			s:        "cat",
			t:        "rat",
			expected: false,
		},
	}

	for _, test := range tests {
		if actual := hashmap.ValidAnagram(test.s, test.t); actual != test.expected {
			t.Errorf("\nS: %v\tT: %v\nExpected: %v\tActual: %v\n", test.s, test.t, test.expected, actual)
		}
	}
}
