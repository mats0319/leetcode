package main

func maxArea(height []int) int {
	var (
		j               = len(height) - 1
		i, area, result int
	)

	for i < j {
		if height[i] < height[j] {
			area = height[i] * (j - i)
			i++
		} else {
			area = height[j] * (j - i)
			j--
		}

		if result < area {
			result = area
		}
	}

	return result
}
