package slidingwindow

import "math"

// Problem: https://leetcode.com/problems/minimum-size-subarray-sum/description

func MinimumSizeSubarraySum(target int, nums []int) int {
	minSize := math.MaxInt
	left, right := 0, 0
	sum := nums[left]

	for {
		if sum >= target {
			if minSize > (right - left + 1) {
				minSize = (right - left + 1)
			}
			if sum == target {
				sum -= nums[left]
				left++
				right++
				if right >= len(nums) {
					break
				}
				sum += nums[right]
			} else {
				sum -= nums[left]
				left++
			}
		} else {
			right++
			if right >= len(nums) {
				break
			}
			sum += nums[right]
		}
	}

	if minSize == math.MaxInt {
		return 0
	}
	return minSize
}
