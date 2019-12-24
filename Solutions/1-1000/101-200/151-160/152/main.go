package mario

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	curMax, curMin, result := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			curMax, curMin = curMin, curMax
		}

		curMax = max(nums[i], curMax*nums[i])
		curMin = min(nums[i], curMin*nums[i])

		result = max(result, curMax)
	}

	return result
}

func max(a, b int) (result int) {
	if a > b {
		result = a
	} else {
		result = b
	}

	return
}

func min(a, b int) (result int) {
	if a < b {
		result = a
	} else {
		result = b
	}

	return
}
