package mario

func searchInsert(nums []int, target int) int {
	index := 0
	for index < len(nums) {
		if target <= nums[index] {
			break
		}

		index++
	}

	return index
}
