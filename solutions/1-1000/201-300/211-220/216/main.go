package mario

var res = make([][]int, 0)
var temp = make([]int, 0)

func combinationSum3(k int, n int) [][]int {
	// for lc
	{
		res = make([][]int, 0)
		temp = make([]int, 0)
	}

	if !hasSolution(n, k) {
		return res
	}

	dfs(0, k, n)

	return res
}

func hasSolution(n, k int) bool {
	max, min := 0, 0
	for i := 1; i <= k; i++ {
		min += i
	}
	for i := 9; i > 9-k; i-- {
		max += i
	}

	return min <= n && n <= max
}

func dfs(used, k, n int) {
	if k == 0 {
		if n == 0 {
			res = append(res, append(make([]int, 0), temp...))
		}
		return
	}

	for i := used + 1; i <= 9; i++ {
		temp = append(temp, i)
		dfs(i, k-1, n-i)
		temp = temp[:len(temp)-1]
	}

	return
}
