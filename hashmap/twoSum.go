package hashmap

func TwoSum(nums []int, target int) []int {
	hash := make(map[int]int)
	for i, num := range nums {
		if j, ok := hash[target-num]; ok {
			return []int{i, j}
		}
		hash[num] = i
	}
	return []int{-1, -1}
}
