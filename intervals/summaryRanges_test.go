package intervals_test

import (
	"slices"
	"testing"

	"github.com/vedantwankhade/katas/leetcode/top-interview-150/intervals"
)

func TestSummaryRanges(t *testing.T) {
	tests := []struct {
		nums     []int
		expected []string
	}{
		{
			nums:     []int{0, 1, 2, 4, 5, 7},
			expected: []string{"0->2", "4->5", "7"},
		},
		{
			nums:     []int{0, 2, 3, 4, 6, 8, 9},
			expected: []string{"0", "2->4", "6", "8->9"},
		},
	}

	for _, test := range tests {
		if actual := intervals.SummaryRanges(test.nums); !slices.Equal(test.expected, actual) {
			t.Errorf("\nNums: %v\nExpected: %v\tActual: %v\n", test.nums, test.expected, actual)
		}
	}
}
