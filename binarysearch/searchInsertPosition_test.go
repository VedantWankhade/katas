package binarysearch_test

import (
	"testing"

	"github.com/vedantwankhade/katas/leetcode/top-interview-150/binarysearch"
)

func TestSearchInsertPosition(t *testing.T) {
	tests := []struct {
		nums     []int
		target   int
		expected int
	}{
		{
			nums:     []int{1, 3, 5, 6},
			target:   5,
			expected: 2,
		},
		{
			nums:     []int{1, 3, 5, 6},
			target:   2,
			expected: 1,
		},
		{
			nums:     []int{1, 3, 5, 6},
			target:   7,
			expected: 4,
		},
	}
	for _, test := range tests {
		if actual := binarysearch.SearchInsertPosition(test.nums, test.target); actual != test.expected {
			t.Errorf("\nNums: %v\tTarget: %v\nExpected: %v\tActual: %v\n", test.nums, test.target, test.expected, actual)
		}
	}
}
