package mario

func minPatches(nums []int, n int) int {
	count := 0

	index := 0
	cover := 1
	for cover <= n {
		if index < len(nums) && nums[index] <= cover {
			cover += nums[index]
			index++
		} else {
			cover *= 2
			count++
		}
	}

	return count
}
