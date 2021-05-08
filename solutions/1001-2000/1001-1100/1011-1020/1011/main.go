package mario

func shipWithinDays(weights []int, D int) int {
	left := 0
	right := 0
	for i := range weights {
		right += weights[i]
		if weights[i] > left {
			left = weights[i]
		}
	}

	maxWeight := 0
	for left <= right {
		maxWeight = left + (right-left)/2

		useDays := 1
		currWeight := 0
		for i := 0; i < len(weights); i++ {
			if currWeight+weights[i] <= maxWeight {
				currWeight += weights[i]
				continue
			}

            useDays++
            if useDays > D {
                break
            }

            currWeight = weights[i]
		}

		if useDays > D {
			left = maxWeight + 1
			maxWeight++ // 防止没有下一次循环时，带出错误答案
		} else {
			right = maxWeight - 1
		}
	}

	return maxWeight
}
