package mario

func searchRange(nums []int, target int) []int {
	left := getLeft(nums, target)
	if left == -1 {
		return []int{-1, -1}
	}

	return []int{left, getRight(nums, target)}
}

func getLeft(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)>>1
		if nums[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}

	if left == len(nums) || nums[left] != target {
		left = -1
	}

	return left
}

func getRight(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)>>1
		if nums[mid] > target {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return right - 1
}
