package mario

func find132pattern(nums []int) bool {
	if len(nums) < 3 {
		return false
	}

	part1 := nums[0]

	isFound := false
ALL:
	for i := 1; i+1 < len(nums); i++ {
		part3 := nums[i] // fixed part3
		if nums[i-1] < part1 {
			part1 = nums[i-1]
		}

		if part1 >= part3 {
			continue
		}

		// loop part2
		for part2Index := i + 1; part2Index < len(nums); part2Index++ {
			if part1 < nums[part2Index] && nums[part2Index] < part3 {
				isFound = true
				break ALL
			}
		}
	}

	return isFound
}
