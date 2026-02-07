package arraystring_test

import (
	"slices"
	"testing"

	"github.com/vedantwankhade/katas/leetcode/top-interview-150/arraystring"
)

func TestMergeSortedArray(t *testing.T) {
	tests := []struct {
		nums1    []int
		m        int
		nums2    []int
		n        int
		expected []int
	}{
		{
			nums1:    []int{1, 2, 3, 0, 0, 0},
			nums2:    []int{2, 5, 6},
			m:        3,
			n:        3,
			expected: []int{1, 2, 2, 3, 5, 6},
		},
		{
			nums1:    []int{1},
			nums2:    []int{},
			m:        1,
			n:        0,
			expected: []int{1},
		},
		{
			nums1:    []int{0},
			m:        0,
			nums2:    []int{1},
			n:        1,
			expected: []int{1},
		},
	}

	for _, test := range tests {
		nums1In := make([]int, len(test.nums1))
		copy(nums1In, test.nums1)
		arraystring.MergeSortedArray(test.nums1, test.m, test.nums2, test.n)
		if !slices.Equal(test.nums1, test.expected) {
			t.Errorf("\nnums1: %v\tnums2: %v\nExpected: %v\tActual: %v", nums1In, test.nums2, test.expected, test.nums1)
		}
	}
}
