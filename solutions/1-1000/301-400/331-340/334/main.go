package mario

func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}

	sequence := [2]int{nums[0]}
	mediumFlag := false
	containsTriplet := false
	for i := 1; i < len(nums); i++ {
		value := nums[i]

		// update small value
		// if value == sequence[0], update is unnecessary, but write to decrease redundancy code
		if value <= sequence[0] {
			sequence[0] = value
			continue
		}

		// update medium value, same reason on "="
		if !mediumFlag || value <= sequence[1] {
			sequence[1] = value
			mediumFlag = true
			continue
		}

		containsTriplet = true
		break
	}

	return containsTriplet
}
