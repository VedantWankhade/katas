package bitmanipulation_test

import (
	"testing"

	bitmanipulation "github.com/vedantwankhade/katas/interviews/coding-problems/primitive/bit_manipulation"
	"github.com/vedantwankhade/katas/interviews/coding-problems/test"
)

func TestParity(t *testing.T) {
	testCases := []test.TestCase[int, int]{
		{Input: 0, Expected: 0},
		{Input: 1, Expected: 1},
		{Input: 2, Expected: 1},
	}
	test.Test(t, bitmanipulation.ParityCountSetBits, testCases)
	test.Test(t, bitmanipulation.ParityDropLowestSetBit, testCases)
}
