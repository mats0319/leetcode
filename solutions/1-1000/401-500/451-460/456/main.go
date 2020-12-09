package mario

func find132pattern(nums []int) bool {
	if len(nums) < 3 {
		return false
	}

	min := make([]int, len(nums))
	min[0] = nums[0]
	// init
	for i := 1; i < len(min); i++ {
		min[i] = small(min[i-1], nums[i])
	}

	// loop 'j'
	isFound := false
	for i := 1; i+1 < len(min) && !isFound; i++ {
		part1 := min[i-1]
		part3 := nums[i]
		index := i+1
		for index < len(min) && !isFound {
			if part1 < nums[index] && nums[index] < part3 {
				isFound = true
			}

			index++
		}
	}

	return isFound
}

func small(a, b int) (res int) {
	if a < b {
		res = a
	} else {
		res = b
	}

	return
}
