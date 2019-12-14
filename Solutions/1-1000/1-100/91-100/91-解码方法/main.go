package mario

func numDecodings(s string) int { // s != ""
	dp := make([]int, len(s)+1)
	dp[0] = 1
	if s[0] != '0' {
		dp[1] = 1
	}
	for i := 2; i < len(s)+1; i++ {
		if s[i-1] == '0' {
			if s[i-2] == '1' || s[i-2] == '2' {
				dp[i] = dp[i-2]
			}
		} else {
			dp[i] = dp[i-1]
			if s[i-2] == '1' || (s[i-2] == '2' && ('1' <= s[i-1] && s[i-1] <= '6')) {
				dp[i] += dp[i-2]
			}
		}
	}

	return dp[len(s)]
}
