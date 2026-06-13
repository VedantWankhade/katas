package binarysearch

func SearchInsertPosition(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	mid := 0

	for lo <= hi {
		mid = lo + (hi-lo)/2
		if nums[mid] < target {
			lo = mid + 1
		} else if nums[mid] > target {
			hi = mid - 1
		} else {
			return mid
		}
	}
	if target < nums[mid] {
		return mid
	}
	return mid + 1
}

