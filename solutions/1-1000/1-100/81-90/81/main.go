package mario

// nums.length >= 1
func search(nums []int, target int) bool {
	isFound := false
	for i := range nums {
		if nums[i] == target {
			isFound = true
			break
		}
	}

	return isFound
}
