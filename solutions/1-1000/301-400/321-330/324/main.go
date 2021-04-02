package mario

import "sort"

func wiggleSort(nums []int) {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	deepCopy := make([]int, len(nums))
	for i := range nums {
		deepCopy[i] = nums[i]
	}

	for i := 0; i < len(nums); i++ {
		if i&1 == 0 {
			nums[i] = deepCopy[(len(deepCopy)+i)/2]
		} else {
			nums[i] = deepCopy[i/2]
		}
	}

	return
}
