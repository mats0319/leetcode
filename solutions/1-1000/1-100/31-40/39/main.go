package mario

func combinationSum(candidates []int, target int) [][]int {
	// 动态规划：f(n) = {n}[if exist n] + {f(n-1)+1}[if exist 1] + {f(n-2)+2}[if exist 2] + ... + {f(n/2)+n-n/2}[if exist n-n/2]
	numberMap := make(map[int]struct{}, len(candidates))
	for i := range candidates {
		numberMap[candidates[i]] = struct{}{}
	}

	dp := make([][][]int, 1, target+1) // deprecated first item

	for n := 1; n <= target; n++ { // use 'n' loop 'target'
		res := make([][]int, 0)
		for i := n - 1; i >= n/2; i-- { // use 'i' loop [n-1, n/2]
			if _, ok := numberMap[n-i]; !ok {
				continue
			}

			for j := range dp[i] {
				item := make([]int, len(dp[i][j]))
				copy(item, dp[i][j])
				item = append(item, n-i)

				if !contains(res, item) {
					res = append(res, item)
				}
			}
		}

		if _, ok := numberMap[n]; ok {
			res = append(res, []int{n})
		}

		dp = append(dp, res)
	}

	return dp[target]
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
