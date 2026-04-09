package arraystring

import "strings"

func LengthOfLastWordUsingBuiltInMethods(s string) int {
	trimmed := strings.TrimSpace(s)
	parts := strings.Split(trimmed, " ")
	return len(parts[len(parts)-1])
}

func LengthOfLastWord(s string) int {
	i := len(s) - 1
	for ; i >= 0 && s[i] == ' '; i-- {

	}

	count := 0
	for ; i >= 0 && s[i] != ' '; i-- {
		count++
	}

	return count
}
