package mario

func combinationSum2(candidates []int, target int) [][]int {
	// only can use once
	// dp: f(n) = {f(0),n}[if n exist] + {f(1),n-1}[if n-1 exist} + ... + {f(n-1),1}[if 1 exist], del duplicated items
	// start: f(0) = make([][]int, 1)
	numMap := make(map[int]int, len(candidates)) // number - appear times
	for i := range candidates {
		numMap[candidates[i]]++
	}

	dp := make([][][]int, target+1) // define on length
	dp[0] = make([][]int, 1)        // set dp[0]

	for i := 1; i <= target; i++ {
		res := make([][]int, 0)

		for n := 0; n < i; n++ {
			if _, ok := numMap[i-n]; !ok {
				continue
			}

			for j := range dp[n] {
				item := make([]int, len(dp[n][j]))
				copy(item, dp[n][j])
				item = append(item, i-n)

				if isValid(item, numMap) && !contains(res, item) {
					res = append(res, item)
				}
			}
		}

		dp[i] = res
	}

	return dp[target]
}

// isValid return if 'm' map can make 'slice'
func isValid(slice []int, m map[int]int) bool {
	sliceMap := make(map[int]int, len(m))
	for i := range slice {
		sliceMap[slice[i]]++
	}

	isValidFlag := true
	for key, value := range sliceMap {
		v, ok := m[key]
		if !ok || value > v {
			isValidFlag = false
			break
		}
	}

	return isValidFlag
}

func contains(slice [][]int, value []int) bool {
	containsFlag := false
	for i := range slice {
		if compareOnIntSlice(slice[i], value) {
			containsFlag = true
			break
		}
	}

	return containsFlag
}

// compareOnIntSlice return if two int slice are equal, ignore items order
func compareOnIntSlice(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	aMap := make(map[int]int, len(a))
	for i := range a {
		aMap[a[i]]++
		aMap[b[i]]--
	}

	isEqualFlag := true
	for _, v := range aMap {
		if v != 0 {
			isEqualFlag = false
			break
		}
	}

	return isEqualFlag
}
