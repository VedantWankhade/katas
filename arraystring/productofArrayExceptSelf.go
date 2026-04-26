package arraystring

// Problem: https://leetcode.com/problems/product-of-array-except-self/description

func ProductOfArrayExceptSelf(nums []int) []int {
	prefixProduct, postfixProduct := make([]int, len(nums)), make([]int, len(nums))

	for i := range nums {
		if i == 0 {
			prefixProduct[i] = 1
		} else {
			prefixProduct[i] = nums[i-1] * prefixProduct[i-1]
		}
	}

	for i := len(nums) - 1; i >= 0; i-- {
		if i == len(nums)-1 {
			postfixProduct[i] = 1
		} else {
			postfixProduct[i] = nums[i+1] * postfixProduct[i+1]
		}
	}

	for i := range nums {
		nums[i] = prefixProduct[i] * postfixProduct[i]
	}

	return nums
}

func ProductOfArrayExceptSelfConstantSpaceKinda(nums []int) []int {
	out := make([]int, len(nums))

	prefix := 1
	for i := range out {
		out[i] = prefix
		prefix *= nums[i]
	}

	postfix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		out[i] *= postfix
		postfix *= nums[i]
	}

	return out
}
