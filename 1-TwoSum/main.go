package main

import "fmt"

func main() {
	n := []int{2, 7, 11, 13}
	fmt.Println(isValid(n, twoSum(n, 9), 9))
}

func isValid(src, is []int, target int) bool {
	var result int
	for i := range is {
		result += src[is[i]]
	}
	return result == target
}

func twoSum(nums []int, target int) []int {
	length := len(nums)
	m := make(map[int]int, length)
	for i := 0; i < length; i++ {
		if v, ok := m[target-nums[i]]; ok {
			return []int{v, i}
		}
		m[nums[i]] = i
	}
	return nil
}
