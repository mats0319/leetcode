package mario

// input is valid
func sortArrayByParityII(A []int) []int {
	res := make([]int, len(A)) // attention: define on length

	index := 0
	for _, v := range A {
		if v%2 == 0 {
			res[index] = v
			index += 2
		}
	}

	index = 1
	for _, v := range A {
		if v%2 == 1 {
			res[index] = v
			index += 2
		}
	}

	return res
}
