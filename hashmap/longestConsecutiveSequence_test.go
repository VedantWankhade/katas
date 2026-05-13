package hashmap_test

import (
	"testing"

	"github.com/vedantwankhade/katas/leetcode/top-interview-150/hashmap"
)

func TestLongestConsecutiveSequence(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{100, 4, 200, 1, 3, 2},
			expected: 4,
		},
		{
			nums:     []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			expected: 9,
		},
		{
			nums:     []int{1, 0, 1, 2},
			expected: 3,
		},
	}

	for _, test := range tests {
		if actual := hashmap.LongestConsecutiveSequence(test.nums); actual != test.expected {
			t.Errorf("\nNums: %v\nExpected: %v\tActual: %v\n", test.nums, test.expected, actual)
		}
	}
}
