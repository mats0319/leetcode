package mario

func moveZeroes(nums []int) {
	value, zero := firstIndexOfValueAndZero(nums)
	if value == -1 || zero == -1 {
		return
	}

	for value < len(nums) {
		nums[value], nums[zero] = nums[zero], nums[value]

		for zero++; zero < len(nums) && nums[zero] != 0; zero++ {
		}

		for value = zero + 1; value < len(nums) && nums[value] == 0; value++ {
		}
	}

	return
}

func firstIndexOfValueAndZero(is []int) (int, int) {
	value, zero := -1, -1

	for i, v := range is {
		if v == 0 {
			zero = i
			break
		}
	}

	for i := zero + 1; i < len(is); i++ {
		if is[i] != 0 {
			value = i
			break
		}
	}

	return value, zero
}
