package primitive

import "testing"

func TestCountActiveBits(t *testing.T) {
	tests := []struct {
		num      uint
		expected uint
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 1},
		{7, 3},
	}

	for _, test := range tests {
		actual := CountActiveBits(test.num)
		if actual != test.expected {
			t.Logf("Test: %d\tExpected: %d\tActual: %d", test.num, test.expected, actual)
		}
	}
}
