package katas

import "fmt"

/*

https://leetcode.com/problems/longest-substring-without-repeating-characters/

Given a string s, find the length of the longest substring

without duplicate characters.



Example 1:

Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3. Note that "bca" and "cab" are also correct answers.

Example 2:

Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.

Example 3:

Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.



Constraints:

    0 <= s.length <= 5 * 104
    s consists of English letters, digits, symbols and spaces.

*/

func LongestSubstringWithoutRepeating(s string) int {
	if len(s) < 2 {
		return len(s)
	}

	inWindow := make(map[rune]int)
	max := 0
	currentLength := 0

	for i, j := 0, 0; i < len(s) && j < len(s); {
		currChar := rune(s[j])
		fmt.Printf("i=%d, j=%d, currChar=%c, currLen=%d, inWindow=%v\n", i, j, currChar, currentLength, inWindow)
		if n, ok := inWindow[currChar]; !ok {
			inWindow[currChar] = j
			j++
		} else {
			if max < currentLength {
				max = currentLength
			}
			currentLength = j - i
			i = n + 1
			delete(inWindow, currChar)
		}
	}

	if max < currentLength {
		max = currentLength
	}

	return max
}
