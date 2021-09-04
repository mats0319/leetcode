package mario

func splitArray(nums []int, m int) int {
	left, right := 0, 0
	for i := range nums {
		if nums[i] > left {
			left = nums[i]
		}

		right += nums[i]
	}

	maxSum := 0
	for left <= right {
		maxSum = left + (right-left)/2

		count := 1
		currSum := 0
		for i := 0; i < len(nums); i++ {
			if currSum+nums[i] <= maxSum {
				currSum += nums[i]
				continue
			}

			count++
			if count > m {
				break
			}
			currSum = nums[i]
		}

		if count > m {
			left = maxSum + 1
			maxSum++
		} else {
			right = maxSum - 1
		}
	}

	return maxSum
}
