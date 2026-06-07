package hashmap_test

import (
	"testing"

	"github.com/vedantwankhade/katas/leetcode/top-interview-150/hashmap"
)

func TestContainsDuplicateII(t *testing.T) {
	tests := []struct {
		nums     []int
		k        int
		expected bool
	}{
		{
			nums:     []int{1, 2, 3, 1},
			k:        3,
			expected: true,
		},
		{
			nums:     []int{1, 0, 1, 1},
			k:        1,
			expected: true,
		},
		{
			nums:     []int{1, 2, 3, 1, 2, 3},
			k:        2,
			expected: false,
		},
	}

	for _, test := range tests {
		if actual := hashmap.ContainsDuplicateII(test.nums, test.k); actual != test.expected {
			t.Errorf("\nNums: %v\tK: %v\nExpected: %v\tActual: %v\n", test.nums, test.k, test.expected, actual)
		}
	}
}
