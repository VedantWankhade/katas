package arraystring

import (
	"bytes"
	"strings"
)

// Problem: https://leetcode.com/problems/reverse-words-in-a-string/description

func ReverseWordsInAString(s string) string {
	s = strings.TrimSpace(s)
	var out bytes.Buffer
	i := len(s) - 1
	j := len(s) - 1
	for i > -1 {
		for j > -1 && s[j] != ' ' {
			j--
		}

		for tempJ := j + 1; tempJ <= i; tempJ++ {
			out.WriteByte(s[tempJ])
		}

		for j > -1 && s[j] == ' ' {
			j--
		}
		if j > -1 {
			out.WriteByte(' ')
		}
		i = j
	}

	return out.String()
}
