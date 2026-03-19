package arraystring_test

import (
	"slices"
	"testing"

	"github.com/vedantwankhade/katas/leetcode/top-interview-150/arraystring"
)

func TestRotateArrayAuxArray(t *testing.T) {
	tests := []struct {
		nums     []int
		k        int
		expected []int
	}{
		{
			nums:     []int{1, 2, 3, 4, 5, 6, 7},
			k:        3,
			expected: []int{5, 6, 7, 1, 2, 3, 4},
		},
	}

	for _, test := range tests {
		inNums := make([]int, len(test.nums))
		copy(inNums, test.nums)
		arraystring.RotateArrayAuxArray(test.nums, test.k)
		if !slices.Equal(test.nums, test.expected) {
			t.Errorf("\nNums: %v\tK: %v\nActual: %v\tExpected: %v", inNums, test.k, test.nums, test.expected)
		}
	}
}

func TestRotateArray(t *testing.T) {
	tests := []struct {
		nums     []int
		k        int
		expected []int
	}{
		{
			nums:     []int{1, 2, 3, 4, 5, 6, 7},
			k:        3,
			expected: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			nums:     []int{-1, -100, 3, 99},
			k:        2,
			expected: []int{3, 99, -1, -100},
		},
		{
			nums:     []int{1},
			k:        1,
			expected: []int{1},
		},
		{
			nums:     []int{1, 2},
			k:        7,
			expected: []int{2, 1},
		},
	}

	for _, test := range tests {
		inNums := make([]int, len(test.nums))
		copy(inNums, test.nums)
		arraystring.RotateArrayReversal(test.nums, test.k)
		if !slices.Equal(test.nums, test.expected) {
			t.Errorf("\nNums: %v\tK: %v\nActual: %v\tExpected: %v", inNums, test.k, test.nums, test.expected)
		}
	}
}
