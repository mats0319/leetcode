package main

func threeSum(nums []int) [][]int {
    if len(nums) < 3 {
        return nil
    }
    
    sort(nums)

    var (
        left, right, t int
        result         = make([][]int, 0)
    )

    for i := 0; i < len(nums)-2; i++ {
        left = i + 1
        right = len(nums) - 1

        for left < right {
            t = nums[i] + nums[left] + nums[right]
            if t == 0 {
                result = append(result, []int{nums[i], nums[left], nums[right]})

                // skip repeated left and right
                for ; left < right && nums[left] == nums[left+1]; left++ {
                }
                for ; left < right && nums[right] == nums[right-1]; right-- {
                }

                left++
                right--
            } else if t > 0 {
                right--
            } else {
                left++
            }
        }

        // skip repeated i
        for i < len(nums)-2 && nums[i] == nums[i+1] {
            i++
        }
    }
    
    return result
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
