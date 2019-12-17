package mario

func isMatch(s string, p string) bool {
    dp := make([][]bool, len(s) + 1)
    for i := range dp {
        dp[i] = make([]bool, len(p) + 1)
    }
    dp[0][0] = true
    for i := 1; i <= len(p) && p[i-1] == '*'; i++ {
        dp[0][i] = true
    }

    for i := 1; i <= len(s); i++ {
        for j := 1; j <= len(p); j++ {
            if p[j-1] == '?' || p[j-1] == s[i-1] {
                dp[i][j] = dp[i-1][j-1]
            } else if p[j-1] == '*' {
                dp[i][j] = dp[i][j-1] || dp[i-1][j]
            }
        }
    }

    return dp[len(s)][len(p)]
}
