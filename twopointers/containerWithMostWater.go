package twopointers

// Problem: https://leetcode.com/problems/container-with-most-water

func ContainerWithMostWater(height []int) int {
	maxArea := 0

	for lo, hi := 0, len(height)-1; lo < hi; {
		minHeight := height[lo]
		if height[lo] > height[hi] {
			minHeight = height[hi]
		}
		area := (hi - lo) * minHeight
		if area > maxArea {
			maxArea = area
		}
		if height[lo] < height[hi] {
			lo++
		} else {
			hi--
		}
	}
	return maxArea
}
