package mario

// nums.length >= 0
// 0 <= nums[i] <= 50
// 0 <= val <= 100
func removeElement(nums []int, val int) int {
	if len(nums) == 0 || val > 50 {
		return len(nums)
	}

	lastIndex := len(nums) - 1
	for i := 0; i <= lastIndex; i++ {
		if nums[i] == val {
			nums[i], nums[lastIndex] = nums[lastIndex], nums[i]
			i--
			lastIndex--
		}
	}

	return lastIndex + 1 // return num is count, not index
}
