package stack

import (
	"fmt"
	"strings"
)

// Problem: https://leetcode.com/problems/simplify-path/

func SimplifyPath(path string) string {
	stack := []string{}
	pathParts := strings.Split(path, "/")

	for _, part := range pathParts {
		if part == "." || part == " " || part == "" {
			continue
		}
		if part == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, part)
		}
	}
	fmt.Printf("%v\t%d\n", stack, len(stack))
	return "/" + strings.Join(stack, "/")
}
