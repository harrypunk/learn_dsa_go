package easy

// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/
func MaxProfit(prices []int) int {
	days := len(prices)
	var cash = make([]int, days)
	var hold = make([]int, days)

	cash[0] = 0
	hold[0] = -prices[0]
	for i := 1; i < days; i++ {
		hold[i] = max(hold[i-1], cash[i-1]-prices[i])
		cash[i] = max(cash[i-1], hold[i-1]+prices[i])
	}

	return cash[days-1]
}

func max(x int, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock
func MaxProfitOnePass(prices []int) int {
	minPrice := prices[0]
	maxProfit := 0

	for i := 0; i < len(prices); i++ {
		var price = prices[i]
		if price < minPrice {
			minPrice = price
		} else if (price - minPrice) > maxProfit {
			maxProfit = price - minPrice
		}
	}
	return maxProfit
}
