package mario

func findTargetSumWays(nums []int, target int) int {
	resMap := make(map[int]int, 0) // res - times
	resMap[nums[0]]++
	resMap[-nums[0]]++

	for i := 1; i < len(nums); i++ {
		newResMap := make(map[int]int, 0)

		for r, times := range resMap {
			newResMap[r+nums[i]] += times
			newResMap[r-nums[i]] += times
		}

		resMap = newResMap
	}

	return resMap[target]
}
