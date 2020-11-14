package mario

// arr1.length, arr2.length <= 1000
// arr2[i] ∈ arr1
func relativeSortArray(arr1 []int, arr2 []int) []int {
	frequency := make([]int, 1001) // attention: define on length

	for _, v := range arr1 {
		frequency[v]++
	}

	index := 0
	for _, v := range arr2 {
		for frequency[v] > 0 {
			arr1[index] = v

			index++
			frequency[v]--
		}
	}

	for i := range frequency {
		for frequency[i] > 0 {
			arr1[index] = i

			index++
			frequency[i]--
		}
	}

	return arr1
}
