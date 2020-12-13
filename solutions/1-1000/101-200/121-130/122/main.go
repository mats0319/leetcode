package mario

func maxProfit(prices []int) int {
    if len(prices) < 2 {
        return 0
    }

    hasStock := false
    buy, sell := 0, 0
    profit := 0
    for i := 1; i < len(prices); i++ {
        if !hasStock {
            if prices[i] <= prices[i-1] {
                buy = i
            } else {
                hasStock = true
                sell = i
            }
        } else {
            if prices[i] >= prices[i-1] {
                sell = i
            } else {
                hasStock = false
                profit += prices[sell] - prices[buy]
                buy = i
            }
        }
    }

    if hasStock {
        profit += prices[sell] - prices[buy]
    }

    return profit
}
