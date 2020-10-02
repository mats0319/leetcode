package mario

import "math"

// input.length >= 3
func minimumOperations(leaves string) int {
    dp := make([][3]int, len(leaves))
    dp[0][0] = isYellow(leaves[0])
    dp[0][1] = math.MaxInt32
    dp[0][2] = math.MaxInt32
    dp[1][2] = math.MaxInt32

    for i := 1; i < len(leaves); i++ {
        dp[i][0] = dp[i-1][0] + isYellow(leaves[i])
        dp[i][1] = min(dp[i-1][0], dp[i-1][1]) + isRed(leaves[i])
        if i > 1 {
            dp[i][2] = min(dp[i-1][1], dp[i-1][2]) + isYellow(leaves[i])
        }
    }

    return dp[len(leaves)-1][2]
}

func min(a, b int) (res int) {
    if a > b {
        res = b
    } else {
        res = a
    }

    return res
}

func isRed(char byte) int {
    res := 0
    if char == 'r' {
        res = 1
    }

    return res
}

func isYellow(char byte) int {
    res := 0
    if char == 'y' {
        res = 1
    }

    return res
}
