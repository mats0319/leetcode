package main

import (
    "math"
)

func threeSumClosest(nums []int, target int) int {
    var (
        left, right, diff, t, sum int
        closest                   = math.MaxInt32
    )

    sort(nums)

    for i := 0; i < len(nums)-2; i++ {
        left = i + 1
        right = len(nums) - 1

        for left < right {
            t = nums[i] + nums[left] + nums[right]
            diff = t - target

            if diff < 0 {
                left++
            } else if diff > 0 {
                right--
            } else {
                sum = target
                return target
            }

            if diff < 0 {
                diff = -diff
            }
            if diff < closest {
                closest = diff
                sum = t
            }
        }
    }

    return sum
}

// quick sort
func sort(is []int) {
    if len(is) <= 1 {
        return
    }

    var small int
    {
        var big int
        for i := 1; i < len(is); i++ {
            if is[i] > is[0] {
                big++
            } else {
                small++
                if big != 0 {
                    is[i], is[small] = is[small], is[i]
                }
            }
        }
        if small != 0 {
            is[0], is[small] = is[small], is[0]
        }
    }

    sort(is[:small])
    sort(is[small+1:])

    return
}
