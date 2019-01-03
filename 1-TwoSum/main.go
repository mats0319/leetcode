package main

import "fmt"

func main() {
	n := []int{2, 7, 11, 13}
	fmt.Println(twoSum(n, 9))
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
