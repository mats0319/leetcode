package mario

func minDistance(word1 string, word2 string) int {
    if len(word1) == 0 || len(word2) == 0 {
        return len(word1) + len(word2)
    }

    dp := make([]int, len(word2)+1)
    for i := range dp {
        dp[i] = i
    }

    var replace, t int
    for i := 1; i <= len(word1); i++ {
        replace = dp[0]
        dp[0] = i
        for j := 1; j <= len(word2); j++ {
            t = dp[j]
            if word1[i-1] == word2[j-1] {
                dp[j] = replace
            } else {
                dp[j] = min(min(dp[j], dp[j-1]), replace)+1
            }
            replace = t
        }
    }

    return dp[len(word2)]
}

func min(a, b int) (result int) {
    if a < b {
        result = a
    } else {
        result = b
    }

    return
}
