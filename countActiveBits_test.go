package katas_test

import (
	"testing"

	katas "github.com/vedantwankhade/katas/interviews/coding-problems"
	"github.com/vedantwankhade/katas/interviews/coding-problems/test"
)

func TestCountActiveBits(t *testing.T) {
	test.Test(t, katas.CountActiveBits, []test.TestCase[uint, uint]{
		{Input: 0, Expected: 0},
		{Input: 1, Expected: 1},
		{Input: 2, Expected: 1},
		{Input: 3, Expected: 2},
		{Input: 4, Expected: 1},
		{Input: 7, Expected: 3},
	})

}
