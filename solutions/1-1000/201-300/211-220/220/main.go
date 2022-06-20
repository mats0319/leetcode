package mario

// |i-j| <= k && |nums[i]-nums[j]| <= t
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if len(nums) < 2 {
		return false
	}

	list := newDoublyList()
	_ = list.addReturn(0, nums[0])

	isValid := false
	for i := 1; i < len(nums); i++ {
		if i > k {
			list.del(i - k - 1)
		}

		node := list.addReturn(i, nums[i])
		if list.validDifference(node, t) {
			isValid = true
			break
		}
	}

	return isValid
}
