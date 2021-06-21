package mario

// nums.length >= 1
func minSubArrayLen(target int, nums []int) int {
	start := 0
	end := 0
	sum := nums[0]
	for end+1 < len(nums) && sum < target {
		end++
		sum += nums[end]
	}

	if sum < target {
		return 0
	}

	length := end - start + 1
	for start < len(nums) {
		sum -= nums[start]
		start++
		for end+1 < len(nums) && sum < target {
			end++
			sum += nums[end]
		}

		if sum >= target && length > end-start+1 {
			length = end - start + 1
		}
	}

	return length
}
