package mario

func trap(height []int) int {
	if len(height) < 3 { // define height is nil
		return 0
	}

	i := 0
	for i < len(height) && height[i] <= 0 {
		i++
	}
	leftMax := height[i]

	rightBig := make([]int, len(height))
	rightBig[len(rightBig)-1] = height[len(height)-1]
	for j := len(rightBig) - 2; j > i; j-- {
		if height[j] > rightBig[j+1] {
			rightBig[j] = height[j]
		} else {
			rightBig[j] = rightBig[j+1]
		}
	}

	res := 0
	for i++; i < len(height)-1; i++ {
		if height[i-1] > leftMax {
			leftMax = height[i-1]
		}

		rightMax := rightBig[i]

		if leftMax > height[i] && rightMax > height[i] {
			res += min(leftMax, rightMax) - height[i]
		}
	}

	return res
}

func min(a, b int) (res int) {
	if a < b {
		res = a
	} else {
		res = b
	}

	return res
}
