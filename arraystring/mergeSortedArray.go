package arraystring

// https://leetcode.com/problems/merge-sorted-array/description

func MergeSortedArray(nums1 []int, m int, nums2 []int, n int) {
	if len(nums1) == 0 || n == 0 {
		return
	}

	left, final, right := m-1, len(nums1)-1, n-1

	for left >= 0 && right >= 0 {
		leftItem, rightItem := nums1[left], nums2[right]
		if leftItem > rightItem {
			nums1[final] = leftItem
			left--
		} else {
			nums1[final] = rightItem
			right--
		}
		final--
	}

	if left < 0 {
		for ; final >= 0; final, right = final-1, right-1 {
			nums1[final] = nums2[right]
		}
	} else {
		for ; final >= 0; final, left = final-1, left-1 {
			nums1[final] = nums1[left]
		}
	}
}
