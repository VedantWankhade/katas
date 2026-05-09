package twopointers

// Problem: https://leetcode.com/problems/is-subsequence/description/

func IsSubsequence(s, t string) bool {
	if len(t) < len(s) {
		return false
	}

	p1, p2 := 0, 0

	for p1 < len(s) && p2 < len(t) && (len(s)-p1) <= (len(t)-p2) {
		if s[p1] == t[p2] {
			p1++
		}
		p2++
	}

	if p1 >= len(s) {
		return true
	}
	return false
}
