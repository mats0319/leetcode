package mario

func maxProfit(prices []int) (result int) {
    if len(prices) < 2 {
        return 0
    }

    var (
        minBid = prices[0]
        profit int
    )
    for i := 1; i < len(prices); i++ {
        if minBid > prices[i-1] {
            minBid = prices[i-1]
        }

        profit = prices[i]-minBid

        if profit > result {
            result = profit
        }
    }

    return
}
