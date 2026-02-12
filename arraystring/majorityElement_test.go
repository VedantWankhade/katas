package arraystring_test

import (
	"testing"

	"github.com/vedantwankhade/katas/leetcode/top-interview-150/arraystring"
)

func TestMajorityElementHashMap(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{3, 2, 3},
			expected: 3,
		},
		{
			nums:     []int{2, 2, 1, 1, 1, 2, 2},
			expected: 2,
		},
		{
			nums:     []int{6, 5, 5},
			expected: 5,
		},
	}

	for _, test := range tests {
		actual := arraystring.MajorityElementHashMap(test.nums)
		if actual != test.expected {
			t.Errorf("\nnums: %v\nexpected: %v\tactual: %v", test.nums, test.expected, actual)
		}
	}
}

func TestMajorityElementBoyerMoore(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{3, 2, 3},
			expected: 3,
		},
		{
			nums:     []int{2, 2, 1, 1, 1, 2, 2},
			expected: 2,
		},
		{
			nums:     []int{6, 5, 5},
			expected: 5,
		},
		{
			nums:     []int{3, 3, 4},
			expected: 3,
		},
	}

	for _, test := range tests {
		actual := arraystring.MajorityElementBoyerMoore(test.nums)
		if actual != test.expected {
			t.Errorf("\nnums: %v\nexpected: %v\tactual: %v", test.nums, test.expected, actual)
		}
	}
}
