package katas_test

import (
	"slices"
	"testing"

	katas "github.com/vedantwankhade/katas/interviews/coding-problems"
	"github.com/vedantwankhade/katas/interviews/coding-problems/test"
)

func TestArrayHashIntersection(t *testing.T) {
	test.TestTwoInputsWithComparator(t, katas.ArrayIntersection1, func(a, e []int) bool {
		if len(a) != len(e) {
			return false
		}
		slices.Sort(a)
		slices.Sort(e)
		return slices.Equal(a, e)
	},
		[]test.TwoInputTestCase[[]int, []int, []int]{
			{[]int{1, 2}, []int{2, 3}, []int{2}},
			{[]int{1, 2, 2, 1}, []int{2, 2}, []int{2, 2}},
			{[]int{4, 9, 5}, []int{9, 4, 9, 8, 4}, []int{9, 4}},
		})
}

func TestArrayHashIntersectionOptimized(t *testing.T) {
	test.TestTwoInputsWithComparator(t, katas.ArrayIntersection2, func(a, e []int) bool {
		if len(a) != len(e) {
			return false
		}
		slices.Sort(a)
		slices.Sort(e)
		return slices.Equal(a, e)
	},
		[]test.TwoInputTestCase[[]int, []int, []int]{
			{[]int{1, 2}, []int{2, 3}, []int{2}},
			{[]int{1, 2, 2, 1}, []int{2, 2}, []int{2, 2}},
			{[]int{4, 9, 5}, []int{9, 4, 9, 8, 4}, []int{9, 4}},
		})
}
