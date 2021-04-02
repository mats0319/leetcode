package mario

func firstMissingPositive(nums []int) int {
    exist := make([]bool, len(nums)+1) // exist[0]: if 1 is already exist
    for i := range nums {
        if 0 < nums[i] && nums[i] <= len(nums) {
            exist[nums[i]-1] = true
        }
    }
    
    res := 1
    for i := 0; i < len(exist) && exist[i]; i++ {
        res++
    }

    return res
}
