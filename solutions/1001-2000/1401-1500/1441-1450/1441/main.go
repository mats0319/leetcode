package mario

func buildArray(target []int, n int) []string {
	res := make([]string, 0, 2*n)
	index := 0
	for i := 1; index < len(target); i++ {
		if i == target[index] {
			index++
			res = append(res, "Push")
		} else {
			res = append(res, "Push", "Pop")
		}
	}

	return res
}
