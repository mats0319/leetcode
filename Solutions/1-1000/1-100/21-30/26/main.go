package mario

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	current := 1
	for scan := 1; scan < len(nums); scan++ {
		if nums[scan] != nums[scan-1] {
			nums[current] = nums[scan]
			current++
		}
	}

	return current
}
