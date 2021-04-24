package mario

import (
    "sort"
    "strconv"
)

// nums.length >= 1
func largestNumber(nums []int) string {
    sort.Slice(nums, func(i, j int) bool {
        x, y := nums[i], nums[j]
        sx := 10
        for sx <= x {
            sx *= 10
        }
        sy := 10
        for sy <= y {
            sy *= 10
        }

        return sy*x+y > sx*y+x
    })

    res := ""
    for i := range nums {
        res += strconv.Itoa(nums[i])
    }

    if res[0] == '0' {
        res = "0"
    }

    return res
}
