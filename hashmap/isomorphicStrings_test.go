package hashmap_test

import (
	"testing"

	"github.com/vedantwankhade/katas/leetcode/top-interview-150/hashmap"
)

func TestIsomorphicStrings(t *testing.T) {
	tests := []struct {
		s        string
		t        string
		expected bool
	}{
		{
			s:        "egg",
			t:        "add",
			expected: true,
		},
		{
			s:        "f11",
			t:        "b24",
			expected: false,
		},
		{
			s:        "paper",
			t:        "title",
			expected: true,
		},
		{
			s:        "badc",
			t:        "baba",
			expected: false,
		},
	}

	for _, test := range tests {
		if actual := hashmap.IsomorphicStrings(test.s, test.t); actual != test.expected {
			t.Errorf("\nS: %v\tT: %v\nExpected: %v\tActual: %v\n", test.s, test.t, test.expected, actual)
		}
	}
}
