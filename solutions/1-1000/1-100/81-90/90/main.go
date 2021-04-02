package mario

// 1 <= nums.length <= 10, -10 <= nums[i] <= 10
func subsetsWithDup(nums []int) [][]int {
	res := make([][]int, 0, 1<<len(nums))
	res = append(res, []int{})
	for i := 0; i < len(nums); i++ {
		length := len(res)
		for j := 0; j < length; j++ {
			arr := deepCopy(res[j], nums[i])

			if !isExist(res, arr) {
				res = append(res, arr)
			}
		}
	}

	return res
}

func isExist(array [][]int, item []int) bool {
	isExistFlag := false
	for i := range array {
		if len(array[i]) == len(item) && isEqual(array[i], item) {
			isExistFlag = true
			break
		}
	}

	return isExistFlag
}

func isEqual(a, b []int) bool {
	used := make([]bool, len(b))
	isEqualFlag := true
	for i := 0; i < len(a); i++ {
		isFound := false
		for j := 0; j < len(used); j++ {
			if a[i] == b[j] && !used[j] {
				isFound = true
				used[j] = true
				break
			}
		}

		if !isFound {
			isEqualFlag = false
			break
		}
	}

	return isEqualFlag
}

func deepCopy(array []int, expand ...int) []int {
	res := make([]int, 0, len(array)+len(expand))
	res = append(array[:0:len(array)], array...)
	res = append(res, expand...)

	return res
}
