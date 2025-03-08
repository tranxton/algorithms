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

func Fibonacci(n int) uint64 {
	var memory = make([]uint64, n+1)

	var fibonacci func(int) uint64
	fibonacci = func(n int) uint64 {
		if n < 2 {
			return uint64(n)
		}

		if memory[n] != 0 {
			return memory[n]
		}

		result := fibonacci(n-1) + fibonacci(n-2)
		memory[n] = result

		return result
	}

	return fibonacci(n)
}

func FindLongestCommonSubstring(target string, words []string) string {
	if len(words) == 0 {
		return ""
	}

	maxLongestCommonSubstring, maxWordIndex := 0, -1

	for i, word := range words {
		wordMaxLongestCommonSubstring := 0
		for targetI := 0; targetI < len(target); targetI++ {
			for wordI := 0; wordI < len(word); wordI++ {
				currentLongestCommonSubstring := 0
				for targetI+currentLongestCommonSubstring < len(target) &&
					wordI+currentLongestCommonSubstring < len(word) &&
					target[targetI+currentLongestCommonSubstring] == word[wordI+currentLongestCommonSubstring] {
					currentLongestCommonSubstring++

					if currentLongestCommonSubstring > wordMaxLongestCommonSubstring {
						wordMaxLongestCommonSubstring = currentLongestCommonSubstring
					}
				}
			}
		}

		if wordMaxLongestCommonSubstring > maxLongestCommonSubstring {
			maxLongestCommonSubstring = wordMaxLongestCommonSubstring
			maxWordIndex = i
		}
	}

	if maxWordIndex == -1 {
		return ""
	}

	return words[maxWordIndex]
}
