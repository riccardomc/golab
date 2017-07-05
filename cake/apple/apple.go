package apple

import "math"

//MaxProfitNaive returns max profit in O(n^2)
func MaxProfitNaive(prices []int) int {
	maxProfit := math.MinInt64
	for i := 0; i < len(prices)-1; i++ {
		for j := i + 1; j < len(prices); j++ {
			distance := prices[j] - prices[i]
			if distance > maxProfit {
				maxProfit = distance
			}
		}
	}
	return maxProfit
}

//MaxProfitGreedy returns max profit in O(n)
func MaxProfitGreedy(prices []int) int {
	if len(prices) < 2 {
		return math.MinInt64
	}

	minPrice := prices[0]
	maxProfit := prices[1] - prices[0]

	for i := 1; i < len(prices); i++ {
		profit := prices[i] - minPrice
		if profit > maxProfit {
			maxProfit = profit
		}

		if prices[i] < minPrice {
			minPrice = prices[i]
		}
	}

	return maxProfit
}
