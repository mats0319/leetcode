package mario

// findPairs find | nums[i] - nums[j] | = k
func findPairs(nums []int, k int) int {
	resMap := make(map[int]struct{})   // 1 - exist, key is smaller one, if 'k' and 'small' are sure, 'big' is sure
	valueMap := make(map[int]struct{}) // 1 - exist, key is value, not index

	for i := range nums {
		if _, ok := valueMap[nums[i]+k]; ok { // nums[j] - nums[i] = k, nums[j] = nums[i]+k
			resMap[nums[i]] = struct{}{}
		}
		if _, ok := valueMap[nums[i]-k]; ok { // nums[i] - nums[j] = k, nums[j] = nums[i]-k
			resMap[nums[i]-k] = struct{}{}
		}

		valueMap[nums[i]] = struct{}{}
	}

	return len(resMap)
}
