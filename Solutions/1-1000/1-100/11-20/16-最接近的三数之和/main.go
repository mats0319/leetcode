package main

import (
	"github.com/mats9693/leetcode/utils/sort"
	"math"
)

func threeSumClosest(nums []int, target int) int {
	var (
		left, right, diff, t, sum int
		closest                   = math.MaxInt32
	)

	sort.QuickSort(nums)

	for i := 0; i < len(nums)-2; i++ {
		left = i + 1
		right = len(nums) - 1

		for left < right {
			t = nums[i] + nums[left] + nums[right]
			diff = t - target

			if diff < 0 {
				left++
			} else if diff > 0 {
				right--
			} else {
				sum = target
				return target
			}

			if diff < 0 {
				diff = -diff
			}
			if diff < closest {
				closest = diff
				sum = t
			}
		}
	}

	return sum
}
