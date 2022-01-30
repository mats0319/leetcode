package mario

func findClosestElements(arr []int, k int, x int) []int {
	left, right := 0, len(arr)-1
	for right-left > 1 {
		mid := left + (right-left)/2
		if arr[mid] == x {
			left = mid
			right = mid
			break
		} else if arr[mid] < x {
			left = mid
		} else {
			right = mid
		}
	}

	if x-arr[left] <= arr[right]-x {
		right = left
	} else {
		left = right
	}

	for right-left+1 < k {
		leftDiff := -1
		if left-1 >= 0 {
			leftDiff = x - arr[left-1]
		}

		rightDiff := -1
		if right+1 < len(arr) {
			rightDiff = arr[right+1] - x
		}

		if leftDiff < 0 {
			right++
		} else if rightDiff < 0 {
			left--
		} else {
			if leftDiff <= rightDiff {
				left--
			} else {
				right++
			}
		}
	}

	return arr[left : right+1]
}
