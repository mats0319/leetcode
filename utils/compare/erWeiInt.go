package cmp

// CompareOnErWeiInt compare if two er-wei int slice is equal
func CompareOnErWeiInt(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	var (
		bMap  = make(map[int][]int, len(b))
		index int
	)

	for i := range b {
		bMap[i] = b[i]
	}

	for i := range a {
		index = contain(bMap, a[i])
		if index != -1 {
			delete(bMap, index)
		} else {
			return false
		}
	}

	return true
}

// if map contain 'is', return map.key
// or return -1 means not contain
func contain(m map[int][]int, is []int) int {
	var (
		isFlag  []int
		isMatch bool
	)

	for i := range m {
		isFlag = make([]int, len(is))

		for j := range m[i] {
			for k := range is {
				if m[i][j] == is[k] && isFlag[k] == 0 { // match and not repeated
					isFlag[k] = 1
					break
				}
			}
		}

		isMatch = true
		for j := range isFlag {
			if isFlag[j] != 1 {
				isMatch = false
				break
			}
		}

		if isMatch {
			return i
		}
	}

	return -1
}
