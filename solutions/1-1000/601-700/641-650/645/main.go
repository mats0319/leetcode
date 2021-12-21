package mario

func findErrorNums(nums []int) []int {
    res := make([]int, 2)
    for i := 0; i < len(nums); i++ {
        targetIndex := nums[i] - 1
        if targetIndex == i {
            continue // number already on it's right position
        }

        if nums[i] == nums[targetIndex] {
            res[0] = nums[i] // duplicate
        } else {
            nums[i], nums[targetIndex] = nums[targetIndex], nums[i]
            i--
        }
    }

    for i := range nums {
        if nums[i] != i+1 {
            res[1] = i + 1
            break
        }
    }

    return res
}
