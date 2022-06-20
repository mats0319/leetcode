package mario

// repeatedNTimes n+1 different number, only one repeated n times, means others are unique
func repeatedNTimes(nums []int) int {
	m := make(map[int]struct{})

	res := 0
	for i := range nums {
		_, ok := m[nums[i]]
		if ok {
			res = nums[i]
			break
		}

		m[nums[i]] = struct{}{}
	}

	return res
}
