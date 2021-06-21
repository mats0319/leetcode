package mario

func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			break
		} else if sum > target {
			right--
		} else {
			left++
		}
	}

	return []int{left + 1, right + 1}
}
