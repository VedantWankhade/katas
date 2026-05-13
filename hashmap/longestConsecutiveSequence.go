package hashmap

// PRoblem: https://leetcode.com/problems/longest-consecutive-sequence/description

// [100,4,200,1,3,2]
func LongestConsecutiveSequence(nums []int) int {
	set := make(map[int]struct{})
	for _, n := range nums {
		set[n] = struct{}{}
	}
	count := make(map[int]int)
	for k := range set {
		if _, ok := count[k]; !ok {
			count[k] = 1 + longestConsecutiveSubsequence(set, count, k-1)
		}
	}
	max := 0
	for _, v := range count {
		if max < v {
			max = v
		}
	}
	return max
}

func longestConsecutiveSubsequence(set map[int]struct{}, count map[int]int, x int) int {
	if _, ok := set[x]; !ok {
		return 0
	}
	if c, ok := count[x]; ok {
		return c
	}
	count[x] = 1 + longestConsecutiveSubsequence(set, count, x-1)
	return count[x]
}
