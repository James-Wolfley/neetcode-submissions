func maxProfit(prices []int) int {
	min := prices[0]
	result := 0

	for _, price := range prices {
		if price < min {
			min = price
		} else if price-min > result {
			result = price - min
		}
	}

	return result
}
