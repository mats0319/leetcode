package mario

func permute(nums []int) [][]int {
    return dfs(make([]int, 0, len(nums)), nums, make([]bool, len(nums)), make([][]int, 0))
}

func dfs(currArray, nums []int, used []bool, res [][]int) [][]int {
    if len(currArray) == len(used) {
        return append(res, append(currArray[:0:0], currArray...))
    }

    for i := 0; i < len(used); i++ {
        if !used[i] {
            used[i] = true

            res = dfs(append(currArray, nums[i]), nums, used, res)

            used[i] = false
        }
    }

    return res
}
