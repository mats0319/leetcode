package mario

func reversePairs(nums []int) int {
	double := make([]int, len(nums))
	for i, v := range nums {
		double[i] = v * 2
	}

	count := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > double[j] {
				count++
			}
		}
	}

	return count
}
