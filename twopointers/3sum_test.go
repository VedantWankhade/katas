package twopointers_test

import (
	"reflect"
	"sort"
	"testing"

	"github.com/vedantwankhade/katas/leetcode/top-interview-150/twopointers"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		nums     []int
		expected [][]int
	}{
		{
			nums:     []int{-1, 0, 1, 2, -1, -4},
			expected: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			nums:     []int{0, 1, 1},
			expected: [][]int{},
		},
		{
			nums:     []int{0, 0, 0},
			expected: [][]int{{0, 0, 0}},
		},
	}

	for _, test := range tests {
		if actual := twopointers.ThreeSum(test.nums); !testEqual(actual, test.expected) {
			t.Errorf("\nNums: %v\nExpected: %v\tActual: %v\n", test.nums, test.expected, actual)
		}
	}
}

func testEqual(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	normalize := func(s [][]int) [][]int {
		res := make([][]int, len(s))

		for i, inner := range s {
			cp := append([]int(nil), inner...)
			sort.Ints(cp)
			res[i] = cp
		}

		sort.Slice(res, func(i, j int) bool {
			for k := 0; k < len(res[i]) && k < len(res[j]); k++ {
				if res[i][k] != res[j][k] {
					return res[i][k] < res[j][k]
				}
			}
			return len(res[i]) < len(res[j])
		})

		return res
	}

	return reflect.DeepEqual(normalize(a), normalize(b))
}
