package mario

func findDuplicate(nums []int) int {
    exist := make([]int, len(nums))
    res := 0
    for i := 0; i < len(nums); i++ {
        exist[nums[i]]++
        if exist[nums[i]] > 1 {
            res = nums[i]
            break
        }
    }

    return res
}
