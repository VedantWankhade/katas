package arraystring_test

import (
	"slices"
	"testing"

	"github.com/vedantwankhade/katas/leetcode/top-interview-150/arraystring"
)

func TestProductOfArrayExceptSelf(t *testing.T) {
	tests := []struct {
		nums     []int
		expected []int
	}{
		{
			nums:     []int{1, 2, 3, 4},
			expected: []int{24, 12, 8, 6},
		},
		{
			nums:     []int{-1, 1, 0, -3, 3},
			expected: []int{0, 0, 9, 0, 0},
		},
	}

	for _, test := range tests {
		actual := arraystring.ProductOfArrayExceptSelf(test.nums)
		if !slices.Equal(actual, test.expected) {
			t.Errorf("\nNums: %v\nExpected: %v\tActual: %v\n", test.nums, test.expected, actual)
		}
	}
}

func TestProductOfArrayExceptSelfCibstabtSpaceKinda(t *testing.T) {
	tests := []struct {
		nums     []int
		expected []int
	}{
		{
			nums:     []int{1, 2, 3, 4},
			expected: []int{24, 12, 8, 6},
		},
		{
			nums:     []int{-1, 1, 0, -3, 3},
			expected: []int{0, 0, 9, 0, 0},
		},
	}

	for _, test := range tests {
		actual := arraystring.ProductOfArrayExceptSelfConstantSpaceKinda(test.nums)
		if !slices.Equal(actual, test.expected) {
			t.Errorf("\nNums: %v\nExpected: %v\tActual: %v\n", test.nums, test.expected, actual)
		}
	}
}
