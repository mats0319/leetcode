package main

import "fmt"

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
	fmt.Println(findMedianSortedArrays([]int{1, 4}, []int{2, 3}))
	fmt.Println(findMedianSortedArrays([]int{1}, []int{2, 3}))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var num1, num2 int

	len1 := len(nums1)
	len2 := len(nums2)

	length := len1 + len2
	flag := length%2 - 1 // flag = -1 or 0.

	for {
		if len1 == 0 {
			num1 = nums2[length/2]
			num2 = nums2[length/2+flag]
			break
		}
		if len2 == 0 {
			num1 = nums1[length/2]
			num2 = nums1[length/2+flag]
			break
		}

		var i, j int = 0, 0
		for c := 0; c < length/2+flag && i < len1 && j < len2; c++ {
			if nums1[i] <= nums2[j] {
				i++
			} else {
				j++
			}
		}

		if i == len1 {
			j = length/2 + flag - len1
			num1 = nums2[j]
			num2 = nums2[j-flag]
			break
		}
		if j == len2 {
			i = length/2 + flag - len2
			num1 = nums1[i]
			num2 = nums1[i-flag]
			break
		}

		if nums1[i] <= nums2[j] {
			num1 = nums1[i]
			i -= flag
		} else {
			num1 = nums2[j]
			j -= flag
		}

		if i == len1 {
			num2 = nums2[j]
			break
		}
		if j == len2 {
			num2 = nums1[i]
			break
		}

		if nums1[i] <= nums2[j] {
			num2 = nums1[i]
		} else {
			num2 = nums2[j]
		}
		break
	}

	return float64(num1+num2) / float64(2)
}
