package mario

// nums is not nil
func canPartition(nums []int) bool {
	target, isValidFlag := isValid(nums)
	if !isValidFlag {
		return isValidFlag
	}

	dp := make([]bool, target+1)
	dp[0] = true
	for i := 0; i < len(nums); i++ {
		v := nums[i]
		for j := target; j >= v; j-- {
			dp[j] = dp[j] || dp[j-v]
		}
	}

	return dp[target]
}

func isValid(nums []int) (int, bool) {
	max := nums[0]
	sum := 0
	for _, n := range nums {
		if n > max {
			max = n
		}

		sum += n
	}

	return sum / 2, sum%2 == 0 && max*2 <= sum
}
