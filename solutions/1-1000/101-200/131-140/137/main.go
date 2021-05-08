package mario

func singleNumber(nums []int) int {
    appear := make(map[int]int)

    for i := 0; i < len(nums); i++ {
        appear[nums[i]]++
    }

    res := 0
    for key, value := range appear {
        if value == 1 {
            res = key
            break
        }
    }

    return res
}
