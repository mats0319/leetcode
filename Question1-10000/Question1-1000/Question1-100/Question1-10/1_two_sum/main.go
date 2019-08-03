package ts

func twoSum(nums []int, target int) []int {
    m := make(map[int]int, len(nums))

    for i := range nums {
        if v, ok := m[target-nums[i]]; ok {
            return []int{v, i}
        }
        m[nums[i]] = i
    }

    return nil
}
