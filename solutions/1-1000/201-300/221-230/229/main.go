package mario

// find num appear more than array.length/3 times
func majorityElement(nums []int) []int {
    m := make(map[int]int, 0)
    for i := range nums {
        m[nums[i]]++
    }

    res := make([]int, 0, 2)
    for n, times := range m {
        if times > len(nums) / 3 {
            res = append(res, n)
        }
    }

    return res
}
