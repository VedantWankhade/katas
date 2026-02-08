package arraystring_test

import (
	"slices"
	"testing"

	"github.com/vedantwankhade/katas/leetcode/top-interview-150/arraystring"
)

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		nums         []int
		val          int
		expectedK    int
		expectedNums []int
	}{
		{
			nums:         []int{3, 2, 2, 3},
			val:          3,
			expectedK:    2,
			expectedNums: []int{2, 2, 3, 3},
		},
		{
			nums:         []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:          2,
			expectedK:    5,
			expectedNums: []int{0, 1, 3, 0, 4},
		},
	}

	for _, test := range tests {
		actualK := arraystring.RemoveElement(test.nums, test.val)
		slices.Sort(test.expectedNums)
		actualNums := test.nums[:actualK]
		slices.Sort(actualNums)
		if actualK != test.expectedK {
			t.Errorf("\nNums: %v\tVal: %v\nActual K: %v\tExpected K: %v\nActual Nums: %v\tExpected Nums: %v", test.nums, test.val, actualK, test.expectedK, actualNums, test.expectedNums)
		}
	}
}
