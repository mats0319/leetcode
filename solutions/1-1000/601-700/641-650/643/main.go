package mario

func findMaxAverage(nums []int, k int) float64 {
	max := 0
	for i := 0; i < k; i++ {
		max += nums[k]
	}

	sum := max
	for i := k; i < len(nums); i++ {
		sum += nums[i] - nums[i-k]
		if max < sum {
			max = sum
		}
	}

	return float64(max) / float64(k)
}
