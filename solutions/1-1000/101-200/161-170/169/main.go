package mario

func majorityElement(nums []int) int {
    m := make(map[int]int, len(nums)) // res - appear times
    for i := range nums {
        m[nums[i]]++
    }

    var res, maxTimes int
    for k, v := range m {
        if maxTimes < v {
            maxTimes = v
            res = k
        }
    }

    return res
}
