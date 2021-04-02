package mario

func subsets(nums []int) [][]int {
	return dfs(make([]int, 0, len(nums)), nums, 0, make([][]int, 0, 1<<len(nums)))
}

func dfs(currArray, nums []int, index int, res [][]int) [][]int {
	if index == len(nums) {
		return append(res, append(currArray[:0:0], currArray...))
	}

	res = dfs(currArray, nums, index+1, res)

	res = dfs(append(currArray, nums[index]), nums, index+1, res)

	return res
}

func subsets_Iteration(nums []int) [][]int {
	res := make([][]int, 0, 1<<len(nums))
	res = append(res, []int{})
	for i := 0; i < len(nums); i++ {
		length := len(res)
		for j := 0; j < length; j++ {
			res = append(res, deepCopy(res[j], nums[i]))
		}
	}

	return res
}

func deepCopy(array []int, expand ...int) []int {
	res := make([]int, 0, len(array)+len(expand))
	res = append(array[:0:len(array)], array...)
	res = append(res, expand...)

	return res
}
