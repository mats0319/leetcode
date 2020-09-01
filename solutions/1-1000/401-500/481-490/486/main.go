package mario

func PredictTheWinner(nums []int) bool {
    length := len(nums)
    dp := make([]int, length)

    for i := length-2; i >= 0; i-- {
        dp[i] = nums[i]
        for j := i+1; j < length; j++ {
            dp[j] = max(nums[i] - dp[j], nums[j] - dp[j-1])
        }
    }

    return dp[len(nums)-1] >= 0
}

func max(a, b int) (res int) {
    if a > b {
        res = a
    } else {
        res = b
    }

    return
}
