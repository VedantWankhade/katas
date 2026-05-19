package twopointers

import (
	"fmt"
	"slices"
)

// Problem: https://leetcode.com/problems/3sum/

func ThreeSum(nums []int) [][]int {
	slices.Sort(nums)
	out := [][]int{}
	debug := [][]int{}

	for p1 := range nums {
		if p1 > 0 && nums[p1] == nums[p1-1] {
			continue
		}
		p2, p3 := p1+1, len(nums)-1
		target := -(nums[p1])
		for p2 < p3 {
			sum := nums[p2] + nums[p3]
			if sum == target {
				out = append(out, []int{nums[p1], nums[p2], nums[p3]})
				debug = append(debug, []int{p1, p2, p3})
				p2++
				for p2 < p3 && nums[p2] == nums[p2-1] {
					p2++
				}
				p3--
				for p3 > p2 && nums[p3] == nums[p3+1] {
					p3--
				}
			} else if sum < target {
				p2++
			} else {
				p3--
			}

		}
	}
	fmt.Println(debug)
	return out
}
