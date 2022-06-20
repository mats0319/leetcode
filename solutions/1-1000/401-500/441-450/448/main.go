package mario

func findDisappearedNumbers(nums []int) []int {
	target := make([]int, len(nums))
	for i := range nums {
		target[nums[i]-1]++
	}

	res := make([]int, 0)
	for i := range target {
		if target[i] == 0 {
			res = append(res, i+1)
		}
	}

	return res
}
