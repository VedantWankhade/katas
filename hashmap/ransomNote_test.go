package hashmap_test

import (
	"testing"

	"github.com/vedantwankhade/katas/leetcode/top-interview-150/hashmap"
)

func TestRansomNote(t *testing.T) {
	tests := []struct {
		ransomNote string
		magazine   string
		expected   bool
	}{
		{
			ransomNote: "a",
			magazine:   "b",
			expected:   false,
		},
		{
			ransomNote: "aa",
			magazine:   "ab",
			expected:   false,
		},
		{
			ransomNote: "aa",
			magazine:   "aab",
			expected:   true,
		},
		{
			ransomNote: "bg",
			magazine:   "efjbdfbdgfjhhaiigfhbaejahgfbbgbjagbddfgdiaigdadhcfcj",
			expected:   true,
		},
	}

	for _, test := range tests {
		if actual := hashmap.RansomNote(test.ransomNote, test.magazine); actual != test.expected {
			t.Errorf("\nRansomNote: %v\tMagazine: %v\nExpected: %v\tActual: %v\n", test.ransomNote, test.magazine, test.expected, actual)
		}
	}
}
