package mario

func fib(n int) int {
	dp := []int{0, 1}

	if n < 2 {
		return dp[n]
	}

	for i := 2; i <= n; i++ {
		dp[i%2] = dp[0] + dp[1]
	}

	return dp[n%2]
}
