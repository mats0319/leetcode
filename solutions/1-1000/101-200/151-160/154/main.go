package mario

// nums.length >= 1, items may duplicate
func findMin(nums []int) int {
    right := len(nums)-1
    for right > 1 && nums[right] == nums[0] {
        right--
    }

    if nums[0] <= nums[right] { // arr is in ascending order or contains only one elem(one or more times)
        return nums[0]
    }

    left := 1
    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] < nums[mid-1] {
            right = mid
            break
        } else if nums[mid] >= nums[0] {
            left = mid + 1
        } else if nums[mid] < nums[0] {
            right = mid - 1
        }
    }

    return nums[right]
}
