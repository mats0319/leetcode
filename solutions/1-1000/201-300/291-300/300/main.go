package mario

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	maxLength := 0
	for i := 0; i < len(nums); i++ {
		length := 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && length < dp[j]+1 {
				length = dp[j] + 1
			}
		}

		dp[i] = length

		if maxLength < length {
			maxLength = length
		}
	}

	return maxLength
}
