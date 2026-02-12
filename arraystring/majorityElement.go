package arraystring

func MajorityElementHashMap(nums []int) int {
	count := make(map[int]int)
	k := len(nums) / 2
	if len(nums) < 2 {
		return nums[0]
	}

	for _, i := range nums {
		itemCount := count[i]
		count[i] = itemCount + 1
		if count[i] > k {
			return i
		}
	}

	return -1
}

func MajorityElementBoyerMoore(nums []int) int {
	x, count := 0, 0

	for _, i := range nums {
		if count == 0 {
			x = i
		}
		if x == i {
			count++
		} else {
			count--
		}
	}

	count = 0

	for _, i := range nums {
		if x == i {
			count++
		}
	}

	if count > len(nums)/2 {
		return x
	}

	return -1
}
