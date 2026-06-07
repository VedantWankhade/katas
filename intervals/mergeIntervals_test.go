package intervals_test

import (
	"testing"

	"github.com/vedantwankhade/katas/leetcode/top-interview-150/intervals"
	"github.com/vedantwankhade/katas/leetcode/top-interview-150/utils"
)

func TestMergeIntervals(t *testing.T) {
	tests := []struct {
		intervals [][]int
		expected  [][]int
	}{
		{
			intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			expected:  [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			intervals: [][]int{{1, 4}, {4, 5}},
			expected:  [][]int{{1, 5}},
		},
		{
			intervals: [][]int{{4, 7}, {1, 4}},
			expected:  [][]int{{1, 7}},
		},
	}

	for _, test := range tests {
		if actual := intervals.MergeIntervals(test.intervals); !utils.SlicesInt2DFuzzyEqual(test.expected, actual) {
			t.Errorf("\nIntervals: %v\nExpected: %v\tActual: %v\n", test.intervals, test.expected, actual)
		}
	}
}
