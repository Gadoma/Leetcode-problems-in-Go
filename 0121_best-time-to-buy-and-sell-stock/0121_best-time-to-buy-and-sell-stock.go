package goleetcode0121

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func maxProfit(prices []int) int {
	minElement := 10001
	maxProfit := 0

	for i := 1; i < len(prices); i++ {
		minElement = minInt(prices[i-1], minElement)
		maxProfit = maxInt(prices[i]-minElement, maxProfit)
	}

	return maxInt(maxProfit, 0)
}
