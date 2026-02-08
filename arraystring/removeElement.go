package arraystring

// https://leetcode.com/problems/remove-element/description

func RemoveElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	left, right, k := 0, len(nums)-1, 0

	for left <= right {
		if nums[left] == val {
			k++
			nums[right], nums[left] = nums[left], nums[right]
			right--
		} else {
			left++
		}
	}

	return len(nums) - k
}
