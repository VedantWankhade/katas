package hashmap_test

import (
	"testing"

	"github.com/vedantwankhade/katas/leetcode/top-interview-150/hashmap"
	"github.com/vedantwankhade/katas/leetcode/top-interview-150/utils"
)

func TestGroupAnagramsFrequencyArray(t *testing.T) {
	tests := []struct {
		strs     []string
		expected [][]string
	}{
		{
			strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			expected: [][]string{
				{"bat"},
				{"nat", "tan"},
				{"ate", "eat", "tea"},
			},
		},
		{
			strs: []string{""},
			expected: [][]string{
				{""},
			},
		},
		{
			strs: []string{"a"},
			expected: [][]string{
				{"a"},
			},
		},
	}
	for _, test := range tests {
		if actual := hashmap.GroupAnagramsFrequencyArray(test.strs); !utils.SlicesString2DFuzzyEqual(actual, test.expected) {
			t.Errorf("\nStrs: %v\nExpected: %v\tActual: %v\n", test.strs, test.expected, actual)
		}
	}
}

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		strs     []string
		expected [][]string
	}{
		{
			strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			expected: [][]string{
				{"bat"},
				{"nat", "tan"},
				{"ate", "eat", "tea"},
			},
		},
		{
			strs: []string{""},
			expected: [][]string{
				{""},
			},
		},
		{
			strs: []string{"a"},
			expected: [][]string{
				{"a"},
			},
		},
	}
	for _, test := range tests {
		if actual := hashmap.GroupAnagrams(test.strs); !utils.SlicesString2DFuzzyEqual(actual, test.expected) {
			t.Errorf("\nStrs: %v\nExpected: %v\tActual: %v\n", test.strs, test.expected, actual)
		}
	}
}
