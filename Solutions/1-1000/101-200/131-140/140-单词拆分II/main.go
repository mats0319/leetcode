package mario

// todo: draft, learn dfs algorithm and tree structure.
func wordBreak(s string, wordDict []string) []string {
	if !canBreak(s, wordDict) {
		return nil
	}

	result := make([][]string, len(s)+1)
	result[0] = []string{""}
	for i := 1; i < len(result); i++ {
		result[i] = make([]string, 0)
	}

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if len(result[j]) != 0 && isContained(wordDict, s[j:i]) {
				for k := range result[j] {
					result[i] = append(result[i], result[j][k]+" "+s[j:i])
				}
			}
		}
	}

	for i := range result[len(s)] {
		result[len(s)][i] = result[len(s)][i][1:]
	}

	return result[len(s)]
}

func canBreak(s string, wordDict []string) bool {
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
