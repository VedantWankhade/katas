package stack

// Problem: https://leetcode.com/problems/valid-parentheses/description

func ValidParantheses(s string) bool {
	pair := map[rune]rune{
		'}': '{',
		')': '(',
		']': '[',
	}
	stack := []rune{}
	for _, r := range s {
		if len(stack)-1 >= 0 && stack[len(stack)-1] == pair[r] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, r)
		}
	}

	if len(stack) == 0 {
		return true
	}
	return false
}
