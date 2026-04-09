package arraystring

// Problem: https://leetcode.com/problems/roman-to-integer/

func RomanToInteger(s string) int {
	num := 0
	rti := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	num += rti[s[len(s)-1]]
	for i := len(s) - 2; i >= 0; i-- {
		if rti[s[i+1]] > rti[s[i]] {
			num -= rti[s[i]]
		} else {
			num += rti[s[i]]
		}
	}

	return num
}
