package mario

var exist map[[8]int]struct{}

// permuteUnique
// 1 <= nums.length <= 8
// -10 <= nums[i] <= 10
func permuteUnique(nums []int) [][]int {
    exist = make(map[[8]int]struct{}) // for leetcode

    return recurse(nums, make([]bool, 8), make([]int, 0), make([][]int, 0))
}

func recurse(nums []int, used []bool, solution []int, res [][]int) [][]int {
    for index := range nums {
        if used[index] {
            continue
        }

        used[index] = true
        solution = append(solution, nums[index])
        if len(solution) < len(nums) { // not finished a solution
            res = recurse(nums, used, solution, res)
        } else if !isExist(solution) {
            res = append(res, deepCopy(solution))
        }

        used[index] = false
        solution = solution[:len(solution)-1]
    }

    return res
}

func isExist(solution []int) bool {
    data := [8]int{}
    for i := range solution {
        data[i] = solution[i]
    }

    _, isExistFlag := exist[data]
    exist[data] = struct{}{}

    return isExistFlag
}

func deepCopy(s []int) []int {
    res := make([]int, len(s))
    for i := range s {
        res[i] = s[i]
    }

    return res
}
