package mario

func maximalRectangle(matrix [][]byte) (result int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	dp := make([]int, len(matrix[0]))

	var s, minHeight int
	for i := 0; i < len(matrix); i++ {
		for j := range matrix[i] {
			dp[j] = (dp[j] + 1) * int(matrix[i][j]-'0')
		}

		for j := range dp {
			if dp[j] > 0 {
				minHeight = dp[j]
				for k := j; k >= 0 && dp[k] > 0; k-- {
					minHeight = min(minHeight, dp[k])
					s = minHeight * (j - k + 1)
					if result < s {
						result = s
					}
				}
			}
		}
	}

	return
}

func min(a, b int) (result int) {
	if a < b {
		result = a
	} else {
		result = b
	}

	return
}
