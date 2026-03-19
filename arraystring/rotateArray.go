package arraystring

func RotateArrayAuxArray(nums []int, k int) {
	k = k % len(nums)
	aux := make([]int, len(nums))
	for i := range nums {
		aux[(i+k)%len(nums)] = nums[i]
	}
	copy(nums, aux)
}

func RotateArrayReversal(nums []int, k int) {
	k = k % len(nums)
	if k == 0 || len(nums) < 2 {
		return
	}

	reverse := func(nums []int, lo, hi int) {
		if hi <= lo {
			return
		}
		for l, h := lo, hi; l < h; l, h = l+1, h-1 {
			nums[l], nums[h] = nums[h], nums[l]
		}
	}

	reverse(nums, 0, len(nums)-k-1)
	reverse(nums, len(nums)-k, len(nums)-1)
	reverse(nums, 0, len(nums)-1)
}
