package mario

// 0 <= nums[i] <= 100
func smallerNumbersThanCurrent(nums []int) []int {
    count := make([]int, 101) // attention: define on length
    for _, v := range nums {
        count[v]++
    }
    for i := 1; i < len(count); i++ {
        count[i] += count[i-1] // count now means sum of [0,n]
    }

    res := make([]int, len(nums)) // attention: define on length
    for i, v := range nums {
        if v > 0 {
            res[i] = count[v-1]
        } else { // if v == 0
            res[i] = 0
        }
    }

    return res
}
