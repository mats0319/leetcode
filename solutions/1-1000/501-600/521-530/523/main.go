package mario

// nums.length > 0
func checkSubarraySum(nums []int, k int) bool {
    modMap := make(map[int]int) // mod - first index
    modMap[0] = -1
    isExist := false
    for i, prefixMod := 0, 0; i < len(nums); i++ {
        prefixMod = (prefixMod + nums[i]) % k

        v, ok := modMap[prefixMod]
        if !ok {
            modMap[prefixMod] = i
        } else if i-v >= 2 {
            isExist = true
            break
        }
    }

    return isExist
}
