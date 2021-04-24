package mario

// matrix.length >= 1
func minFallingPathSum(matrix [][]int) int {
    dp := matrix[0]
    for i := 1; i < len(matrix); i++ {
        nextLine := matrix[i]
        for j := 0; j < len(dp); j++ {
            minValue := dp[j]
            if j-1 >= 0 && minValue > dp[j-1] {
                minValue = dp[j-1]
            }
            if j+1 < len(matrix[i]) && minValue > dp[j+1] {
                minValue = dp[j+1]
            }

            nextLine[j] += minValue
        }

        dp = nextLine
    }

    return min(dp)
}

// make sure arr is not nil
func min(arr []int) int {
    res := arr[0]
    for i := 1; i < len(arr); i++ {
        if res > arr[i] {
            res = arr[i]
        }
    }

    return res
}
