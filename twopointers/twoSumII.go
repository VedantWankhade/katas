package twopointers

// Problem : https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/description

func TwoSumII(numbers []int, target int) []int {
	for left, right := 0, len(numbers)-1; left < right; {
		currSum := numbers[left] + numbers[right]
		if target == currSum {
			return []int{left + 1, right + 1}
		} else if target < currSum {
			right--
		} else {
			left++
		}
	}
	return []int{-1, -1}
}
