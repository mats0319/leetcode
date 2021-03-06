package mario

func containsDuplicate(nums []int) bool {
	exist := make(map[int]bool, len(nums))

	contains := false
	for _, v := range nums {
		if _, ok := exist[v]; ok {
			contains = true
			break
		} else {
			exist[v] = true
		}
	}

	return contains
}
