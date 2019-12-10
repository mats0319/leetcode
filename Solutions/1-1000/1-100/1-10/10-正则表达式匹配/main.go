package mario

func isMatch(s string, p string) bool {
	isMatched := true
	for {
		if p == "" {
			isMatched = s == ""
			break
		}
		isMatched = reMatch(s, p)
		break
	}
	return isMatched
}

func reMatch(s, p string) (isMatched bool) {
	for {
		se := s == ""
		if p == "" {
			isMatched = se
			break
		}
		for ip := range p {
			if ip < len(p)-1 && p[ip+1] == '*' {
				if !se && (s[ip] == p[ip] || p[ip] == '.') {
					isMatched = reMatch(s[1:], p) || reMatch(s, p[2:])
				} else {
					isMatched = reMatch(s, p[2:])
				}
				break
			}
			if !se && (s[ip] == p[ip] || p[ip] == '.') {
				isMatched = reMatch(s[1:], p[1:])
			}
			break
		}
		break
	}
	return
}

// dynamic programming: todo: space optimize?
func isMatchDP(s, p string) bool {
	lens, lenp := len(s), len(p)
	dp := make([][]bool, lens+1)
	for i := range dp {
		dp[i] = make([]bool, lenp+1)
	}

	dp[0][0] = true // s == p == "" ; dp[0][1] = false, default, omit

	for j := 2; j <= lenp; j++ {
		dp[0][j] = dp[0][j-2] && p[j-1] == '*' // s: "a", p: "b*c*d*a"
	}

	for i := 1; i <= lens; i++ { // i: ith
		for j := 1; j <= lenp; j++ { // j: jth
			if p[j-1] == '*' {
				dp[i][j] = dp[i][j-2] || dp[i-1][j] && (p[j-2] == s[i-1] || p[j-2] == '.') // skill only on bool type
			} else if p[j-1] == s[i-1] || p[j-1] == '.' {
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}

	return dp[lens][lenp]
}
