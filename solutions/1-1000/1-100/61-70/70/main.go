package mario

func climbStairs(n int) int {
	dp := []int{2, 1}
	for i := 3; i <= n; i++ {
		dp[i%2] = dp[0] + dp[1]
	}

	return dp[n%2]
}
