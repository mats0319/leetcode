package mario

func countRangeSum(nums []int, lower int, upper int) int {
    if len(nums) < 1 {
        return 0
    }

    preSum := make([]int, len(nums)) // attention: define on length, [0, i]
    preSum[0] = nums[0]
    for i := 1; i < len(preSum); i++ {
        preSum[i] = preSum[i-1] + nums[i]
    }

    count := 0

    for i := 0; i < len(preSum); i++ {
        preSumI := 0
        if i > 0 {
            preSumI = preSum[i-1]
        }

        for j := i; j < len(preSum); j++ {
            value := preSum[j] - preSumI // summary[i, j]
            if lower <= value && value <= upper {
                count++
            }
        }
    }

    return count
}
