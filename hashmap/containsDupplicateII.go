package hashmap

import "math"

func ContainsDuplicateII(nums []int, k int) bool {
	hash := make(map[int]int)

	for i, n := range nums {
		if m, ok := hash[n]; !ok {
			hash[n] = i
		} else {
			if int(math.Abs(float64(i-m))) <= k {
				return true
			}
			hash[n] = i
		}
	}
	return false
}
