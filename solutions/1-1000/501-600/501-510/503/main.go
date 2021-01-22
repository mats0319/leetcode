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
