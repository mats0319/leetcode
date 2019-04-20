package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}) == 49)
}

func maxArea(height []int) int {
	var (
		maxArea int
		is      = 0
		ie      = len(height) - 1
	)

	for is != ie {
		t := 0
		if height[is] >= height[ie] {
			t = height[ie] * (ie - is)
			ie--
		} else {
			t = height[is] * (ie - is)
			is++
		}
		if t > maxArea {
			maxArea = t
		}
	}
	return maxArea
}
