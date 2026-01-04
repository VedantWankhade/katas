package katas_test

import (
	"testing"

	katas "github.com/vedantwankhade/katas/interviews/coding-problems"
	"github.com/vedantwankhade/katas/interviews/coding-problems/test"
)

func TestParity(t *testing.T) {
	testCases := []test.TestCase[int, int]{
		{Input: 0, Expected: 0},
		{Input: 1, Expected: 1},
		{Input: 2, Expected: 1},
	}
	test.Test(t, katas.ParityCountSetBits, testCases)
	test.Test(t, katas.ParityDropLowestSetBit, testCases)
}
