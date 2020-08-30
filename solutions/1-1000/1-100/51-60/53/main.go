package mario

func maxSubArray(nums []int) (result int) {
	var sum int

	result = nums[0]
	for i := range nums {
		sum += nums[i]
		result = max(sum, result)
		if sum < 0 {
			sum = 0
		}
	}

	return
}

func max(a, b int) (result int) {
	if a > b {
		result = a
	} else {
		result = b
	}

	return
}
