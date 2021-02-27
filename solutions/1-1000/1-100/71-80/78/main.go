package mario

func subsets(nums []int) [][]int {
    return dfs(make([]int, 0, len(nums)), nums, 0, make([][]int, 0, 1 << len(nums)))
}

func dfs(currArray, nums []int, index int, res [][]int) [][]int {
    if index == len(nums) {
        return append(res, append(currArray[:0:0], currArray...))
    }

    res = dfs(currArray, nums, index+1, res)

    res = dfs(append(currArray, nums[index]), nums, index+1, res)

    return res
}
