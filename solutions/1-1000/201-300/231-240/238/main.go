package mario

// nums.length > 1
func productExceptSelf(nums []int) []int {
	left := make([]int, len(nums)) // left[i]: mul[:i), not contains nums[i]
	left[0] = 1
	for i := 1; i < len(left); i++ {
		left[i] = left[i-1] * nums[i-1]
	}

	right := make([]int, len(nums)) // right[i]: mul(i:], not contains nums[i]
	right[len(right)-1] = 1
	for i := len(right)-2; i >= 0; i-- {
		right[i] = right[i+1] * nums[i+1]
	}

	res := make([]int, len(nums))
	for i := 0; i < len(res); i++ {
		res[i] = left[i] * right[i]
	}

	return res
}

func productExceptSelf2(nums []int) []int {
	left := make([]int, len(nums)) // left[i]: mul[:i), not contains nums[i]
	left[0] = 1
	for i := 1; i < len(left); i++ {
		left[i] = left[i-1] * nums[i-1]
	}

	mul := 1
	for i := len(left)-1; i >= 0; i-- {
		left[i] *= mul
		mul *= nums[i]
	}

	return left
}
