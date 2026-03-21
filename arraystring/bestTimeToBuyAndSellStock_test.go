package arraystring_test

import (
	"testing"

	"github.com/vedantwankhade/katas/leetcode/top-interview-150/arraystring"
)

func TestBestTimeToBuyAndSellStockBruteForce(t *testing.T) {
	tests := []struct {
		prices    []int
		expProfit int
	}{
		{
			prices:    []int{7, 1, 5, 3, 6, 4},
			expProfit: 5,
		},
		{
			prices:    []int{7, 6, 4, 3, 1},
			expProfit: 0,
		},
	}

	for _, test := range tests {
		prof := arraystring.BestTimeToBuyAndSellStockBruteForce(test.prices)
		if prof != test.expProfit {
			t.Errorf("Prices: %v\nExpected: %v\tActual: %v\n", test.prices, test.expProfit, prof)
		}
	}
}

func TestBestTimeToBuyAndSellStockSlidingWindow(t *testing.T) {
	tests := []struct {
		prices    []int
		expProfit int
	}{
		{
			prices:    []int{7, 1, 5, 3, 6, 4},
			expProfit: 5,
		},
		{
			prices:    []int{7, 6, 4, 3, 1},
			expProfit: 0,
		},
	}

	for _, test := range tests {
		prof := arraystring.BestTimeToBuyAndSellStockSlidingWindow(test.prices)
		if prof != test.expProfit {
			t.Errorf("Prices: %v\nExpected: %v\tActual: %v\n", test.prices, test.expProfit, prof)
		}
	}
}
