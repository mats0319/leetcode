package mario

func longestValidParentheses(s string) (result int) {
	if len(s) < 2 {
		return 0
	}

	dp := make([]int, len(s)+1)
	for i := 2; i <= len(s); i++ {
		if s[i-1] == ')' {
			if s[i-2] == '(' {
				dp[i] = dp[i-2] + 2
			} else {
				if dp[i-1] > 0 && (i-dp[i-1]-2 >= 0 && s[i-dp[i-1]-2] == '(') {
					dp[i] = dp[i-1] + 2
					index := i - dp[i-1] - 2
					for index >= 0 && dp[index] > 0 {
						dp[i] += dp[index]
						index -= dp[index]
					}
				}
			}

			if result < dp[i] {
				result = dp[i]
			}
		}
	}

	return
}

func longestValidParenthesesWithConstantStorage(s string) (result int) {
	if len(s) < 2 {
		return 0
	}

	var left, right int
	for i := range s {
		if s[i] == '(' {
			left++
		} else {
			right++
		}

		if left == right {
			if result < left+right {
				result = left + right
			}
		} else if left < right {
			left, right = 0, 0
		}
	}

	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}

		if left == right {
			if result < left+right {
				result = left + right
			}
		} else if left > right {
			left, right = 0, 0
		}
	}

	return
}
