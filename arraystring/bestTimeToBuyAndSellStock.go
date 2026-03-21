package arraystring

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description

func BestTimeToBuyAndSellStockBruteForce(prices []int) int {
	maxProfit := 0

	for i := 0; i < len(prices); i++ {
		buy := prices[i]
		for j := i; j < len(prices); j++ {
			profit := prices[j] - buy
			if profit > maxProfit {
				maxProfit = profit
			}
		}
	}
	return maxProfit
}

func BestTimeToBuyAndSellStockSlidingWindow(prices []int) int {
	maxProfit := 0
	minBuy := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < minBuy {
			minBuy = prices[i]
		} else if prices[i]-minBuy > maxProfit {
			maxProfit = prices[i] - minBuy
		}
	}

	return maxProfit
}
