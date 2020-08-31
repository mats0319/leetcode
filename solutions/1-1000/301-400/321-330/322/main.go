package mario

// todo: optimize, only 12.36% time and 17.37% space!
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1) // dp[0] = 0, dp[1]: amount=1

	for i := 1; i < len(dp); i++ {
		alternatives := make([]int, 0)

		for j := range coins {
			if i-coins[j] >= 0 && dp[i-coins[j]] != -1 {
				alternatives = append(alternatives, dp[i-coins[j]])
			}
		}

		if len(alternatives) == 0 {
			dp[i] = -1
			continue
		}

		dp[i] = alternatives[0]
		for j := range alternatives {
			if alternatives[j] < dp[i] {
				dp[i] = alternatives[j]
			}
		}

		dp[i]++
	}

	return dp[amount]
}
