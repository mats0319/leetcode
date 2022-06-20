package mario

func search(nums []int, target int) int {
	index := -1
	for left, right := 0, len(nums)-1; left <= right; {
		mid := left + (right-left)/2

		if mid < 0 || mid >= len(nums) {
			break
		}

		if nums[mid] == target {
			index = mid
			break
		} else if nums[mid] < target {
			left = mid + 1
		} else { // nums[mid] > target
			right = mid - 1
		}
	}

	return index
}
