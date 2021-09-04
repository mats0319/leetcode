package mario

// matrix is always not nil
func searchMatrix(matrix [][]int, target int) bool {
	lineLength := len(matrix[0])

	lineIndex := 0
	isFound := false
	{
		left := 0
		right := len(matrix) - 1
		for left <= right {
			mid := left + (right-left)/2
			if target == matrix[mid][0] || target == matrix[mid][lineLength-1] {
				isFound = true
				break
			} else if matrix[mid][0] < target && target < matrix[mid][lineLength-1] {
				lineIndex = mid
				break
			} else if target < matrix[mid][0] {
				right = mid - 1
			} else { // if target > matrix[mid][lineLength-1]
				left = mid + 1
			}
		}
	}

	if isFound {
		return true
	}

	{
		left := 0
		right := lineLength - 1
		for left <= right {
			mid := left + (right-left)/2
			value := matrix[lineIndex][mid]
			if target == value {
				isFound = true
				break
			} else if target > value {
				left = mid + 1
			} else { // if target < value
				right = mid - 1
			}
		}
	}

	return isFound
}
