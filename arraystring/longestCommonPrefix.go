package arraystring

import "bytes"

// Problem: https://leetcode.com/problems/longest-common-prefix

func longestCommonPrefix(strs []string) string {
	ptr := 0
	out := bytes.Buffer{}
	for ptr < len(strs[0]) {
		c := strs[0][ptr]
		match := true
		for _, s := range strs[1:] {
			if ptr >= len(s) || s[ptr] != c {
				match = false
			}
		}
		if match {
			out.WriteByte(c)
			ptr++
		} else {
			return out.String()
		}
	}
	return out.String()
}
