package slidingwindow_test

import (
	"testing"

	"github.com/vedantwankhade/katas/leetcode/top-interview-150/slidingwindow"
)

func Test_minimumSizeSubarraySum(t *testing.T) {
	tests := []struct {
		target   int
		nums     []int
		expected int
	}{
		{
			target:   7,
			nums:     []int{2, 3, 1, 2, 4, 3},
			expected: 2,
		},
		{
			target:   4,
			nums:     []int{1, 4, 4},
			expected: 1,
		},
		{
			target:   11,
			nums:     []int{1, 1, 1, 1, 1, 1, 1, 1},
			expected: 0,
		},
		{
			target:   11,
			nums:     []int{1, 2, 3, 4, 5},
			expected: 3,
		},
	}

	for _, test := range tests {
		if actual := slidingwindow.MinimumSizeSubarraySum(test.target, test.nums); actual != test.expected {
			t.Errorf("\nTarget: %v\tNums: %v\nExpected: %v\tActual: %v\n", test.target, test.nums, test.expected, actual)
		}
	}
}
