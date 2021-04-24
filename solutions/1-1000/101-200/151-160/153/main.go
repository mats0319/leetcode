package mario

// nums.length >= 1, every elem is unique
func findMin(nums []int) int {
	if nums[0] <= nums[len(nums)-1] { // arr is in ascending order or contains only one elem
		return nums[0]
	}

	index := 0
	{
		left := 1
		right := len(nums) - 1
		for left <= right {
			mid := left + (right-left)/2
			if nums[mid] < nums[mid-1] {
				index = mid
				break
			} else if nums[mid] > nums[0] {
				left = mid + 1
			} else if nums[mid] < nums[0] {
				right = mid - 1
			}
		}
	}

	return nums[index]
}
