package arraystring

// Problem: https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/description

func FindTheIndexOfTheFirstOccurrenceInAString(haystack string, needle string) int {
	if len(needle) > len(haystack) {
		return -1
	}

	i, j := 0, 0

	for i < len(haystack) && j < len(needle) {
		if haystack[i] == needle[j] {
			i++
			j++
		} else {
			i = i - j + 1
			j = 0
		}
	}
	if j == len(needle) {
		return i - len(needle)
	}
	return -1
}
