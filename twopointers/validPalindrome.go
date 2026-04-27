package twopointers

import (
	"strings"
)

// Problem: https://leetcode.com/problems/valid-palindrome

func ValidPalindrome(s string) bool {
	s = strings.ToLower(s)

	for lo, hi := 0, len(s)-1; lo < hi; {
		if isAlphaNumeric(s[lo]) && isAlphaNumeric(s[hi]) {
			if s[lo] != s[hi] {
				return false
			}
			lo++
			hi--
		} else {
			if !isAlphaNumeric(s[lo]) {
				lo++
			}
			if !isAlphaNumeric(s[hi]) {
				hi--
			}
		}
	}
	return true
}

func isAlphaNumeric(b byte) bool {
	const (
		firstValid    = byte(97)
		lastValid     = byte(97 + 25)
		firstValidNum = byte(48)
		lastValidNum  = byte(48 + 9)
	)

	return (b >= firstValid && b <= lastValid) || (b >= firstValidNum && b <= lastValidNum)
}
