package mario

import (
    "fmt"
    "strconv"
)

// nums.length >= 0
func summaryRanges(nums []int) []string {
    if len(nums) < 1 {
        return nil
    }

    nums = append(nums, nums[0])

    res := make([]string, 0)
    traversal := []int{nums[0], nums[0]}
    for i := 1; i < len(nums); i++ {
        if traversal[1] +1 == nums[i] {
            traversal[1]++
            continue
        }

        if traversal[0] == traversal[1] {
            res = append(res, strconv.Itoa(traversal[0]))
        } else {
            res = append(res, fmt.Sprintf("%d->%d", traversal[0], traversal[1]))
        }

        traversal = []int{nums[i], nums[i]}
    }

    return res
}
