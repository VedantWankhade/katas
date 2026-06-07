package intervals

import (
	"slices"
)

func MergeIntervals(intervals [][]int) [][]int {
	for i := range intervals {
		slices.Sort(intervals[i])
	}
	slices.SortFunc(intervals, func(arr1, arr2 []int) int {
		return arr1[0] - arr2[0]
	})

	out := [][]int{}

	out = append(out, intervals[0])
	p1, p2 := 0, 1

	for p2 < len(intervals) {
		left := out[p1]
		right := intervals[p2]

		if left[0] == right[0] {
			out[p1] = []int{left[0], max(left[1], right[1])}
			p2++
			continue
		}

		if right[0] > left[0] && right[0] <= left[1] {
			out[p1] = []int{min(left[0], right[0]), max(left[1], right[1])}
			p2++
			continue
		}

		out = append(out, right)
		p1++
		p2++
	}

	return out
}
