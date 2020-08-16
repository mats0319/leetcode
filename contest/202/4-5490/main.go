package mario

func minDays(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 || n == 3 {
		return 2
	}

	dp := make([]int, n)
	dp[0] = 1
	dp[1] = 2
	dp[2] = 2

	for i := 3; i < len(dp); i++ {
		mod2 := (i+1)%2 == 0
		mod3 := (i+1)%3 == 0

		if mod2 && mod3 {
			dp[i] = min(dp[i-1], dp[i/2], dp[i/3]) + 1
		} else if mod2 {
			dp[i] = min(dp[i-1], dp[i/2]) + 1
		} else if mod3 {
			dp[i] = min(dp[i-1], dp[i/3]) + 1
		} else {
			dp[i] = dp[i-1] + 1
		}
	}

	return dp[n-1]
}

func min(slice ...int) (res int) {
	res = slice[0]
	for i := range slice {
		if slice[i] < res {
			res = slice[i]
		}
	}

	return
}
