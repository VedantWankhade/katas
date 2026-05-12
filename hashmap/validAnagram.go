package hashmap

// Problem: https://leetcode.com/problems/valid-anagram

func ValidAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	hash := make(map[rune]int)
	for _, c := range s {
		if hc, ok := hash[c]; ok {
			hash[c] = hc + 1
		} else {
			hash[c] = 1
		}
	}
	for _, c := range t {
		if hc, ok := hash[c]; ok {
			hash[c] = hc - 1
		} else {
			hash[c] = -1
		}
	}

	for _, v := range hash {
		if v != 0 {
			return false
		}
	}

	return true
}
