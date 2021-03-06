package mario

func wiggleMaxLength(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	res := 1
	prevSub := nums[1] - nums[0]
	if prevSub != 0 {
		res++
	}

	for i := 2; i < len(nums); i++ {
		sub := nums[i] - nums[i-1]
		if (sub > 0 && prevSub <= 0) || (sub < 0 && prevSub >= 0) {
			res++
			prevSub = sub
		}
	}

	return res
}
