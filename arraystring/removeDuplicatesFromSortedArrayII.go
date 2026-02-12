package arraystring

// https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/description

func RemoveDuplicatesFromSortedArrayII(nums []int) int {
	if len(nums) < 3 {
		return len(nums)
	}

	write, current := 2, 2

	for current < len(nums) {
		if nums[current] != nums[write-2] {
			nums[write] = nums[current]
			write++
		}
		current++
	}

	return write
}
