package leet_code

func MaxProfit(prices []int) int {
	minBuyPrice, maxProfit := prices[0], 0

	for i := 1; i < len(prices); i++ {
		price := prices[i]

		if price < minBuyPrice {
			minBuyPrice = price

			continue
		}

		if profit := price - minBuyPrice; profit > maxProfit {
			maxProfit = profit
		}
	}

	if maxProfit > 0 {
		return maxProfit
	}

	return 0
}
