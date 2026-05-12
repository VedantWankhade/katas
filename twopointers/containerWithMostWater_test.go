package twopointers_test

import (
	"testing"

	"github.com/vedantwankhade/katas/leetcode/top-interview-150/twopointers"
)

func TestContainerWithMostWater(t *testing.T) {
	tests := []struct {
		height   []int
		expected int
	}{
		{
			height:   []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			expected: 49,
		},
		{
			height:   []int{1, 1},
			expected: 1,
		},
	}

	for _, test := range tests {
		if actual := twopointers.ContainerWithMostWater(test.height); actual != test.expected {
			t.Errorf("\nHeight: %v\nExpected: %v\tActual: %v", test.height, test.expected, actual)
		}
	}
}
