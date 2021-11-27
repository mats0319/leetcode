package mario

func findDuplicates(nums []int) []int {
	target := make([]int, len(nums))
	res := make([]int, 0)
	for i := range nums {
		target[nums[i]-1]++
		if target[nums[i]-1] == 2 {
			res = append(res, nums[i])
		}
	}

	return res
}
