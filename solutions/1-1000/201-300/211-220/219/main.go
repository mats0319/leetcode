package mario

func containsNearbyDuplicate(nums []int, k int) bool {
	if k > len(nums) {
		k = len(nums)
	}

	exist := make(map[int]struct{}, k+1)
	hasDuplicate := false
	for i := 0; i < len(nums); i++ {
		if len(exist) > k {
			delete(exist, nums[i-k-1])
		}

		if _, ok := exist[nums[i]]; ok {
			hasDuplicate = true
			break
		}

		exist[nums[i]] = struct{}{}
	}

	return hasDuplicate
}
