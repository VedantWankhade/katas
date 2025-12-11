package bitmanipulation_test

import (
	"testing"

	bitmanipulation "github.com/vedantwankhade/katas/interviews/coding-problems/primitive/bit_manipulation"
	"github.com/vedantwankhade/katas/interviews/coding-problems/test"
)

func TestDropLowestSetBit(t *testing.T) {
	test.Test(t, bitmanipulation.DropLowestSetBit, []test.TestCase[int, int]{
		{
			Input:    0, // 0000
			Expected: 0, // 0000
		},
		{
			Input:    1, // 0001
			Expected: 0, // 0000
		},
		{
			Input:    2, // 0010
			Expected: 0, // 0000
		},
		{
			Input:    3, // 0011
			Expected: 2, // 0010
		},
		{
			Input:    6, // 0110
			Expected: 4, // 0100
		},
	})
}
