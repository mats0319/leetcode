package mario

func maxProfit(prices []int) int {
    if len(prices) < 2 {
        return 0
    }

    // dp[i][0]：持有股票；dp[i][1]：不持有股票，因为今天刚卖；dp[i][2]：不持有股票，当日无操作
    dp := []int{-prices[0], 0, 0}
    for i := 1; i < len(prices); i++ {
        profit := make([]int, 3)
        profit[0] = max(dp[0], dp[2]-prices[i])
        profit[1] = dp[0]+prices[i]
        profit[2] = max(dp[2], dp[1])

        dp = profit
    }

    return max(dp[1], dp[2])
}

func max(a, b int) (res int) {
    if a > b {
        res = a
    } else {
        res = b
    }

    return res
}
