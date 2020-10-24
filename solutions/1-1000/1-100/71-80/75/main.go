package mario

func sortColors(nums []int)  {
    count := [3]int{}
    for i := range nums {
        count[nums[i]]++
    }

    index := 0
    for i, c := range count {
        for j := 0; j < c; j++ {
            nums[index] = i
            index++
        }
    }

    return
}
