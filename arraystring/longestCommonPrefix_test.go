package arraystring

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	tests := []struct {
		strs     []string
		expected string
	}{
		{
			strs:     []string{"flower", "flow", "flight"},
			expected: "fl",
		},
		{
			strs:     []string{"dog", "racecar", "car"},
			expected: "",
		},
	}

	for _, test := range tests {
		actual := longestCommonPrefix(test.strs)
		if actual != test.expected {
			t.Errorf("\nStrs: %v\nExpected: %v\tActual: %v", test.strs, test.expected, actual)
		}
	}
}
