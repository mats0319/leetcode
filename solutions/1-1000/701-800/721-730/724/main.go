package mario

func pivotIndex(nums []int) int {
    if len(nums) < 1 {
        return -1
    }
    
    prefixSum := make([]int, len(nums))
    prefixSum[0] = nums[0]
    for i := 1; i < len(prefixSum); i++ {
        prefixSum[i] = prefixSum[i-1]+nums[i]
    }

    mid := -1
    for i := 0; i < len(prefixSum); i++ {
        if prefixSum[i]-nums[i] == prefixSum[len(prefixSum)-1]-prefixSum[i] {
            mid = i
            break
        }
    }

    return mid
}
