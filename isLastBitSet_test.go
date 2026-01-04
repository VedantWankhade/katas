package katas_test

import (
	"testing"

	katas "github.com/vedantwankhade/katas/interviews/coding-problems"
	"github.com/vedantwankhade/katas/interviews/coding-problems/test"
)

func TestIsLastBitSet(t *testing.T) {
	test.Test(t, katas.IsLastBitSet, []test.TestCase[int, bool]{
		{Input: 0, Expected: false},
		{Input: 1, Expected: true},
		{Input: 2, Expected: false},
		{Input: 3, Expected: true},
		{Input: 6, Expected: false},
	})
}
