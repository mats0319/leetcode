package mario

// nums.length >= 1
func search(nums []int, target int) int {
	if nums[0] <= nums[len(nums)-1] { // array move on index: 0 or only contains one elem
		return isExist(nums, target)
	}

	minIndex := firstSmallIndex(nums)
	index := -1
	if nums[0] <= target && target <= nums[minIndex-1] {
		index = isExist(nums[:minIndex], target)
	}
	if nums[minIndex] <= target && target <= nums[len(nums)-1] {
		index = isExist(nums[minIndex:], target)
		if index >= 0 {
			index += minIndex
		}
	}

	return index
}

// firstSmallIndex find first elem small than array[0] and return it's index
func firstSmallIndex(array []int) int {
	index := -1
	for left, right := 0, len(array)-1; left <= right; {
		mid := left + (right-left)/2
		if mid > 0 && array[mid] < array[0] && array[0] <= array[mid-1] {
			index = mid
			break
		}
		if array[mid] >= array[0] {
			left = mid + 1
		} else if array[mid] < array[0] {
			right = mid - 1
		}
	}

	return index
}

// isExist return index
// array is monotonous increase and it's elem is unique
func isExist(array []int, target int) int {
	index := -1
	if len(array) == 1 {
		if array[0] == target {
			index = 0
		}
		return index
	}

	for left, right := 0, len(array)-1; left <= right; {
		mid := left + (right-left)/2
		if array[mid] == target {
			index = mid
			break
		} else if array[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return index
}
