package mario

func nextPermutation(nums []int)  {
    if len(nums) < 2 {
        return
    }

    small := len(nums)-2
    for small >= 0 && nums[small] >= nums[small+1] {
        small--
    }

    if small >= 0 {
        big := len(nums)-1
        for big > small && nums[big] <= nums[small] {
            big--
        }

        nums[small], nums[big] = nums[big], nums[small]
    }

    for l, r := small+1, len(nums)-1; l < r; l, r = l+1, r-1 {
        nums[l], nums[r] = nums[r], nums[l]
    }

    return
}
