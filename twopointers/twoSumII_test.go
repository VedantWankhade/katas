package twopointers_test

import (
	"testing"

	"github.com/vedantwankhade/katas/leetcode/top-interview-150/twopointers"
)

func TestTwoSumII(t *testing.T) {
	tests := []struct {
		nums     []int
		target   int
		expected []int
	}{
		{
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{1, 2},
		},
		{
			nums:     []int{2, 3, 4},
			target:   6,
			expected: []int{1, 3},
		},
		{
			nums:     []int{-1, 0},
			target:   -1,
			expected: []int{1, 2},
		},
	}

	testEqual := func(s1, s2 []int) bool {
		hash := make(map[int]int)
		for _, num := range s1 {
			if count, ok := hash[num]; ok {
				hash[num] = count + 1
			} else {
				hash[num] = 1
			}
		}
		for _, num := range s2 {
			if count, ok := hash[num]; ok {
				hash[num] = count - 1
			} else {
				hash[num] = -1
			}
		}

		for _, v := range hash {
			if v != 0 {
				return false
			}
		}
		return true
	}

	for _, test := range tests {
		if actual := twopointers.TwoSumII(test.nums, test.target); !testEqual(actual, test.expected) {
			t.Errorf("\nNums: %v\tTarget: %v\nExpected: %v\tActual: %v\n", test.nums, test.target, test.expected, actual)
		}
	}
}
