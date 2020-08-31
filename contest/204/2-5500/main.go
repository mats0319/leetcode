package mario

// nums.length >= 1
func getMaxLen(nums []int) int {
	numsSplitBy0 := make([][]int, 0)
	start := 0
	for i := range nums {
		if nums[i] == 0 {
			numsSplitBy0 = append(numsSplitBy0, nums[start:i])
			start = i + 1
		}
	}

	numsSplitBy0 = append(numsSplitBy0, nums[start:])

	max := 0
	for _, subSlice := range numsSplitBy0 {
		posCount := 0
		maxPosCount := 0
		negIndex := make([]int, 0)
		for i := range subSlice {
			if subSlice[i] > 0 {
				posCount++
			} else { // if subSlice[i] < 0, subSlice not contains '0'
				if posCount > maxPosCount {
					maxPosCount = posCount
				}
				posCount = 0
				negIndex = append(negIndex, i)
			}
		}

		if posCount > maxPosCount {
			maxPosCount = posCount
		}

		subMax := 0
		if len(negIndex)%2 == 0 { // 0 neg case contains
			subMax = len(subSlice)
		} else if len(negIndex) == 1 {
			subMax = maxPosCount
		} else {
			subMax = big(len(subSlice)-negIndex[0]-1, negIndex[len(negIndex)-1])
		}

		if subMax > max {
			max = subMax
		}
	}

	return max
}

func big(a, b int) (res int) {
	if a > b {
		res = a
	} else {
		res = b
	}

	return
}
