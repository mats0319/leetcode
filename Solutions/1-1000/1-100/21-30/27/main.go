package mario

func removeElement(nums []int, val int) int {
    if len(nums) == 0 {
        return 0
    }

    lastIndex := len(nums)-1
    for i := 0; i < lastIndex; i++ {
        if nums[i] == val {
            nums[i], nums[lastIndex] = nums[lastIndex], nums[i]
            i--
            lastIndex--
        }
    }

    if nums[lastIndex] != val {
        lastIndex++
    }

    return lastIndex // return num is count, not index
}
