package mario

func findLatestStep(arr []int, m int) int {
	if len(arr) == 1 {
		return 1
	} else if len(arr) == m {
		return m
	}

	str := make([]int, len(arr)+1)
	str[0] = 1

	res := -1
	for i := len(arr) - 1; i > 0; i-- {
		str[arr[i]] = 1

		count := 0
		for j := 1; j < len(str); j++ {
			if str[j] == 0 {
				count++
			} else {
				if count == m {
					break
				}

				count = 0
			}
		}

		if count == m {
			res = i
			break
		}
	}

	return res
}
