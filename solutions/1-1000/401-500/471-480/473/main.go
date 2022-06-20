package mario

import "sort"

func makesquare(matchsticks []int) bool {
	if len(matchsticks) < 4 {
		return false
	}

	sum := 0
	for i := range matchsticks {
		sum += matchsticks[i]
	}

	if sum%4 != 0 {
		return false
	}

	sort.Ints(matchsticks)

	side := sum / 4

	if matchsticks[len(matchsticks)-1] > side {
		return false
	}

	return recursion([4]int{}, matchsticks, len(matchsticks)-1, side)
}

func recursion(fourSide [4]int, options []int, optionIndex int, side int) bool {
	if optionIndex < 0 {
		return true
	}

	isValid := false
	for i := 0; i < 4; i++ {
		if fourSide[i]+options[optionIndex] <= side {
			fourSide[i] += options[optionIndex]

			isValid = recursion(fourSide, options, optionIndex-1, side)
			if isValid {
				break
			}

			fourSide[i] -= options[optionIndex]
		}
	}

	return isValid
}
