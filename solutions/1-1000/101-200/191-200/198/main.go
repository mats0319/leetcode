package mario

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}

	// dp[i]：街道上有i户人家
	dp := []int{0, nums[0]}
	for i := 2; i <= len(nums); i++ {
		dp[i%2] = max(dp[1-(i%2)], dp[i%2]+nums[i-1]) // space optimize ^_^
	}

	return dp[len(nums)%2]
}

func max(a, b int) (result int) {
	if a > b {
		result = a
	} else {
		result = b
	}

	return
}
