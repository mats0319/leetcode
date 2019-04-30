package main

import (
	"fmt"
	"math"
)

// sort before is faster than force do it.
func main() {
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1) == 2)
	fmt.Println(threeSumClosest([]int{0, 1, 2}, 3) == 3)
	fmt.Println(threeSumClosest([]int{1, 1, 1, 0}, -100) == 2)
	fmt.Println(threeSumClosest([]int{1, 1, -1, -1, 3}, -1) == -1)
	fmt.Println(threeSumClosest([]int{-3, -2, -5, 3, -4}, -1) == -2)
	fmt.Println(threeSumClosest([]int{0, 2, 1, -3}, 1) == 0)
}

func threeSumClosest(nums []int, target int) int {
	var (
		left, right, diff, s, sum int
		closest                   = math.MaxInt32
	)

	sort(nums)

match:
	for i := 0; i < len(nums)-2; i++ {
		left = i + 1
		right = len(nums) - 1
		for left < right {
			s = nums[i] + nums[left] + nums[right]
			diff = s - target
			if diff < 0 {
				left++
			} else if diff > 0 {
				right--
			} else {
				sum = target
				break match
			}
			if diff < 0 {
				diff = -diff
			}
			if diff < closest {
				closest = diff
				sum = s
			}
		}
	}
	return sum
}

func sort(is []int) {
	var maxIndex, minIndex int
	for i := range is {
		if is[i] > is[maxIndex] {
			maxIndex = i
		} else if is[i] < is[minIndex] {
			minIndex = i
		}
	}

	var (
		offset = is[minIndex]
		temp   = make([]int, is[maxIndex]-offset+1)
	)
	for i := range is {
		temp[is[i]-is[minIndex]]++
	}
	
	var in int
	for i := range temp {
		for ; temp[i] > 0; temp[i]-- {
			is[in] = i + offset
			in++
		}
	}
}
