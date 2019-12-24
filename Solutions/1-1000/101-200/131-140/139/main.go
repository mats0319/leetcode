package mario

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && isContained(wordDict, s[j:i]) {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]
}

func isContained(strs []string, s string) (result bool) {
	for i := range strs {
		if s == strs[i] {
			result = true
			break
		}
	}

	return
}
