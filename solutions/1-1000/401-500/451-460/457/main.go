package mario

// 1 <= nums.length <= 5000
// -1000 <= nums[i] <= 1000, nums[i] != 0
func circularArrayLoop(nums []int) bool {
	hasLoop := false

	length := len(nums)
	for i := 0; i < length; i++ {
		start := i
		next := ((start+nums[start])%length + length) % length
		if start == next {
			continue
		}

		reached := make([]int, 0, length)
		for isSameDirection(nums[start], nums[next]) && next > start && !contains(reached, next) {
			reached = append(reached, next)
			next = ((next+nums[next])%length + length) % length
		}

		if start == next {
			hasLoop = true
			break
		}
	}

	return hasLoop
}

func isSameDirection(a, b int) bool {
	return (a > 0 && b > 0) || (a < 0 && b < 0)
}

func contains(slice []int, value int) bool {
	containsValue := false
	for i := range slice {
		if slice[i] == value {
			containsValue = true
			break
		}
	}

	return containsValue
}
