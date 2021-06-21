package mario

func findMaxConsecutiveOnes(nums []int) int {
	max := 0
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			count++
		} else {
			max = big(max, count)
			count = 0
		}
	}

	return big(max, count)
}

func big(a, b int) (res int) {
	if a > b {
		res = a
	} else {
		res = b
	}

	return
}
