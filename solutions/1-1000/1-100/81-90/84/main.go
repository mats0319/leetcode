package mario

func largestRectangleArea(heights []int) int {
	max := 0
	for i, v := range heights {
		sum := v * getIntervalLength(heights, i)
		if sum > max {
			max = sum
		}
	}

	return max
}

func getIntervalLength(slice []int, index int) int {
	count := 1
	for i := index-1; i >= 0 && slice[i] >= slice[index]; i-- {
		count++
	}
	for i := index+1; i < len(slice) && slice[i] >= slice[index]; i++ {
		count++
	}

	return count
}
