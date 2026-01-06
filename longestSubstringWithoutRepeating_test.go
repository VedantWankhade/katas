package katas_test

import (
	"testing"

	katas "github.com/vedantwankhade/katas/interviews/coding-problems"
	"github.com/vedantwankhade/katas/interviews/coding-problems/test"
)

func TestLongestSubstringWithoutRepeating(t *testing.T) {
	testCases := []test.TestCase[string, int]{
		{Input: "abcabcbb", Expected: 3},
		{Input: "bbbbb", Expected: 1},
		{Input: "pwwkew", Expected: 3},
		{Input: "dvdf", Expected: 3},
		{Input: "tmmzuxt", Expected: 5},
	}
	test.Test(t, katas.LongestSubstringWithoutRepeating, testCases)
}
