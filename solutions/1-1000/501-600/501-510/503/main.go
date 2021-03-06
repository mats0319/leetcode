package mario

func nextGreaterElements(nums []int) []int {
	index := make([]int, 0, len(nums))
	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		res[i] = -1

		for len(index) > 0 && nums[i] > nums[index[len(index)-1]] {
			res[index[len(index)-1]] = nums[i]
			index = index[:len(index)-1]
		}

		index = append(index, i)
	}

	for i := 0; i < len(nums) && len(index) > 0; i++ {
		for len(index) > 0 && nums[i] > nums[index[len(index)-1]] {
			res[index[len(index)-1]] = nums[i]
			index = index[:len(index)-1]
		}
	}

	return res
}

func nextGreaterElements_DoubleTraversal(nums []int) []int {
	length := len(nums)
	nums = append(nums, nums...)

	res := make([]int, length)
	for i := 0; i < length; i++ {
		res[i] = -1

		for j := i + 1; j < i+length; j++ {
			if nums[j] > nums[i] {
				res[i] = nums[j]
				break
			}
		}
	}

	return res
}
