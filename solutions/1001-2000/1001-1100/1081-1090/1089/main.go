package mario

func duplicateZeros(arr []int) {
	arrCopy := make([]int, 0, len(arr))
	for i := range arr {
		arrCopy = append(arrCopy, arr[i])
	}

	index := 0
	repeatFlag := false
	for i := 0; i < len(arr); i++ {
		if repeatFlag {
			repeatFlag = false
			arr[i] = 0
			continue
		}

		arr[i] = arrCopy[index]

		repeatFlag = arrCopy[index] == 0
		index++
	}
}
