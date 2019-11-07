package main

func findMedianSortedArrays(nums1, nums2 []int) float64 {
	length := len(nums1) + len(nums2)
	flag := length%2 - 1 // flag = 0 (odd) or -1 (even)

	k1 := (length + 1) / 2
	k2 := (length+1)/2 - flag

	return float64(getKthSmallNumber(nums1, nums2, k1)+getKthSmallNumber(nums1, nums2, k2)) / float64(2)
}

// algorithm: if a[k/2] < b[k/2], a[0] ... a[k/2-1] not the k-th small number.
// each time pass k/2 number, time complexity is O(log(m+n))
// by the way, tail recurse make space complexity is O(1)
func getKthSmallNumber(is1, is2 []int, k int) int {
	if len(is1) == 0 {
		return is2[k-1]
	}
	if len(is2) == 0 {
		return is1[k-1]
	}
	if k == 1 {
		return small(is1[0], is2[0])
	}

	index1, index2 := k/2-1, k/2-1
	if len(is1)-1 < k/2-1 {
		index1 = len(is1) - 1
	}
	if len(is2)-1 < k/2-1 {
		index2 = len(is2) - 1
	}

	if is1[index1] < is2[index2] {
		if index1 == 0 {
			index1++
		}
		k -= index1
		is1 = is1[index1:]
	} else {
		if index2 == 0 {
			index2++
		}
		k -= index2
		is2 = is2[index2:]
	}

	return getKthSmallNumber(is1, is2, k)
}

func small(a, b int) (r int) {
	if a < b {
		r = a
	} else {
		r = b
	}

	return
}

//O(m+n), try to use two-divide method to make it O(log(m+n)), which named the k-th small number algorithm.
func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	var num1, num2 int

	len1 := len(nums1)
	len2 := len(nums2)

	length := len1 + len2
	flag := length%2 - 1 // flag = -1 or 0, ignore length is odd or even.

	for {
		// if one array is nil.
		{
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
		}

		// structure: before-num1 num1 num2
		// loop till before-num1 or one array is empty
		var i, j int
		for c := 0; c < length/2+flag && i < len1 && j < len2; c++ {
			if nums1[i] <= nums2[j] {
				i++
			} else {
				j++
			}
		}

		// if one array is empty
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
