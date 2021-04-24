package mario

func removeDuplicates(nums []int) int {
    index := 0
    count := 1
    for i := 1; i < len(nums); i++ {
        if nums[i] == nums[index] {
            count++
            if count <= 2 {
                index++
            }
        } else {
            index++
            count = 1
        }

        nums[index] = nums[i]
    }

    return index+1
}
