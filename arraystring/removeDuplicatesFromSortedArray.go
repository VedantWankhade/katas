package arraystring

// https://leetcode.com/problems/remove-duplicates-from-sorted-array/description

func RemoveDuplicatesFromSortedArray(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	left, right, k := 0, 1, 0

	for right < len(nums)-k {
		if nums[left] == nums[right] {
			for j := right + 1; j < len(nums)-k; j++ {
				nums[j-1] = nums[j]
			}
			k++
		} else {
			left++
			right++
		}
	}

	return len(nums) - k
}

func RemoveDuplicatesFromSortedArrayFast(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	left, right := 0, 1

	for right < len(nums) {
		if nums[left] == nums[right] {
			right++
		} else {
			nums[left+1], nums[right] = nums[right], nums[left+1]
			left++
			right++
		}
	}

	return left + 1
}
