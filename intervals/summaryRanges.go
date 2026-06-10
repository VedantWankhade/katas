package intervals

import (
	"fmt"
	"strconv"
)

// Problem: https://leetcode.com/problems/summary-ranges/description/

func SummaryRanges(nums []int) []string {
	out := []string{}

	if len(nums) == 0 {
		return out
	}

	if len(nums) == 1 {
		return append(out, strconv.Itoa(nums[0]))
	}

	lo, hi := 0, 1

	for ; hi < len(nums); hi++ {
		if nums[hi] != nums[hi-1]+1 {
			if lo == hi-1 {
				out = append(out, fmt.Sprintf("%d", nums[lo]))
			} else {
				out = append(out, fmt.Sprintf("%d->%d", nums[lo], nums[hi-1]))
			}
			lo = hi
		}
	}

	if lo == hi-1 {
		out = append(out, fmt.Sprintf("%d", nums[lo]))
	} else {
		out = append(out, fmt.Sprintf("%d->%d", nums[lo], nums[hi-1]))
	}

	return out
}
