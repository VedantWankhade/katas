package hashmap

// Prblem: https://leetcode.com/problems/isomorphic-strings/description

func IsomorphicStrings(s, t string) bool {
	shash := make(map[byte]byte, len(s))
	thash := make(map[byte]byte, len(s))

	for sp, tp := 0, 0; sp < len(s) && tp < len(t); sp, tp = sp+1, tp+1 {
		if spc, ok := shash[s[sp]]; ok {
			if spc != t[tp] {
				return false
			}
		} else {
			if _, ok := thash[t[tp]]; !ok {
				shash[s[sp]] = t[tp]
				thash[t[tp]] = s[sp]
			} else {
				return false
			}
		}
	}
	return true
}
