package mario

import "sort"

// nums.length >= 1
func largestDivisibleSubset(nums []int) []int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	maxLength := make([]int, len(nums))
	maxLength[0] = 1
	lastIndex := make([]int, len(nums))
	lastIndex[0] = -1
	for i := 1; i < len(nums); i++ {
		index := getLastIndex(nums, maxLength, i)
		if index < 0 {
			maxLength[i] = 1
			lastIndex[i] = -1
		} else {
			maxLength[i] = maxLength[index] + 1
			lastIndex[i] = index
		}
	}

	index := 0
	{
		max := maxLength[0]
		for i := 1; i < len(maxLength); i++ {
			if maxLength[i] > max {
				max = maxLength[i]
				index = i
			}
		}
	}

	res := make([]int, 0, maxLength[index])
	for index >= 0 {
		res = append(res, nums[index])
		index = lastIndex[index]
	}

	return res
}

func getLastIndex(array, length []int, targetIndex int) int {
	index := -1
	{
		maxLength := -1
		for i := 0; i < targetIndex; i++ {
			if length[i] > maxLength && array[targetIndex]%array[i] == 0 {
				maxLength = length[i]
				index = i
			}
		}
	}

	return index
}
