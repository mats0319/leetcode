package mario

func missingNumber(nums []int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		res ^= i+1
		res ^= nums[i]
	}

	return res
}
