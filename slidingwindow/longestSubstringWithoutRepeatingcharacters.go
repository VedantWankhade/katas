package slidingwindow

// Problem: https://leetcode.com/problems/longest-substring-without-repeating-characters/

func LongestSubstringWithoutRepeatingCharacters(s string) int {
	if len(s) < 2 {
		return len(s)
	}

	hash := make(map[byte]int)

	hash[s[0]] = 0
	p := 0
	longest := 1
	long := 1

	for i := 1; i < len(s); {
		current := s[i]
		if existingIndex, ok := hash[current]; ok && existingIndex >= p {
			p = existingIndex + 1
			hash[current] = i
			if i == len(s)-1 {
				return max(longest, long)
			}
			i++
			if long > longest {
				longest = long
			}
			long = i - p
		} else {
			hash[current] = i
			long++
			i++
		}
	}

	return max(longest, long)
}
