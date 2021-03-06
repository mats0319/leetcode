package mario

// s.length = t.length
func equalSubstring(s string, t string, maxCost int) int {
	cost := make([]int, len(s))
	for i := range s {
		cost[i] = calcCost(s[i], t[i])
	}

	maxLength := 0
	for i := 0; i+maxLength < len(cost)+1; i++ {
		length := 0
		costRemain := maxCost
		for j := i; j < len(cost) && costRemain >= cost[j]; j++ {
			costRemain -= cost[j]
			length++
		}

		if maxLength < length {
			maxLength = length
		}
	}

	return maxLength
}

func calcCost(a, b uint8) (res int) {
	if a > b {
		res = int(a - b)
	} else {
		res = int(b - a)
	}

	return
}
