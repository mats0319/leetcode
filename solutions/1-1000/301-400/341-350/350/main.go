package mario

func intersect(nums1 []int, nums2 []int) []int {
	m := make(map[int]int, len(nums1))
	for _, v := range nums1 {
		m[v]++
	}

	res := make([]int, 0, min(len(nums1), len(nums2)))
	for _, v := range nums2 {
		if times, ok := m[v]; ok && times > 0 {
			m[v]--
			res = append(res, v)
		}
	}

	return res
}

func min(a, b int) (res int) {
	if a < b {
		res = a
	} else {
		res = b
	}

	return
}
