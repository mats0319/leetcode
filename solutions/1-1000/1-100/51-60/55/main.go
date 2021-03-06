package mario

func canJump(nums []int) bool {
	maxDistance := nums[0]
	start := 0
	for i := start; i < len(nums) && i <= maxDistance; i++ {
		if maxDistance < i+nums[i] {
			maxDistance = i + nums[i]
		}
	}

	return maxDistance >= len(nums)-1
}
