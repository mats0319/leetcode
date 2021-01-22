package mario

func maxProfit(prices []int, fee int) int {
	if len(prices) < 2 {
		return 0
	}

	profit := []int{0, -prices[0]} // 0: no stock, 1: has stock
	for i := 1; i < len(prices); i++ {
		newProfit := make([]int, 2)
		newProfit[0] = max(profit[0], profit[1]+prices[i]-fee)
		newProfit[1] = max(profit[1], profit[0]-prices[i])

		profit = newProfit
	}

	return profit[0]
}

func max(a, b int) (res int) {
	if a > b {
		res = a
	} else {
		res = b
	}

	return
}
