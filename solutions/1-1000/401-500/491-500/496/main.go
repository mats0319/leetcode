package mario

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	if len(nums1) < 1 || len(nums2) < 1 {
		return nil
	}

	m := make(map[int]int, len(nums2))
	stack := make([]int, 0, len(nums2))
	for i := 0; i < len(nums2); i++ {
		for len(stack) > 0 && nums2[i] > stack[len(stack)-1] {
			m[stack[len(stack)-1]] = nums2[i]
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, nums2[i])
	}

	for len(stack) > 0 {
		m[stack[0]] = -1
		stack = stack[1:]
	}

	res := make([]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		res[i] = m[nums1[i]]
	}

	return res
}
