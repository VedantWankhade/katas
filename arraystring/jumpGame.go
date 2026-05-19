package arraystring

// Problem: https://leetcode.com/problems/jump-game

func JumpGameDP(nums []int) bool {
	dp := make([]bool, len(nums))
	dp[len(nums)-1] = true

	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] == 0 {
			dp[i] = false
		} else {
			furthest := i + nums[i]
			dp[i] = false
			for j := i + 1; j <= furthest; j++ {
				if dp[j] {
					dp[i] = true
					break
				}
			}
		}
	}
	return dp[0]
}

func JumpGameGreedy(nums []int) bool {
	maxReachSoFar := 0
	for i, num := range nums {
		if i > maxReachSoFar {
			return false
		}
		if i+num > maxReachSoFar {
			maxReachSoFar = i + num
		}
	}

	return true
}
