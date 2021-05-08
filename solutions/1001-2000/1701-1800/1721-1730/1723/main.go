package mario

import "sort"

// 1 <= k <= jobs.length <= 12
func minimumTimeRequired(jobs []int, k int) int {
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i] > jobs[j]
	})

	left, right := getMaxAndSum(jobs)
	for left < right {
		mid := left + (right-left)/2

		if isValid(jobs, k, mid) {
			right = mid
		} else {
			left = mid+1
		}
	}

	return left
}

func isValid(arr []int, count, limit int) bool {
	assign := make([]bool, len(arr))
	assigned := 0
	for i := 0; i < count && assigned < len(assign); i++ {
		sum := 0
		for j := 0; j < len(assign); j++ {
			if !assign[j] && sum+arr[j] <= limit {
				assign[j] = true
				assigned++
				sum += arr[j]
			}
		}
	}

	return assigned >= len(assign)
}

func getMaxAndSum(arr []int) (max int, avg int) {
	for i := range arr {
		if arr[i] > max {
			max = arr[i]
		}

		avg += arr[i]
	}

	return
}
