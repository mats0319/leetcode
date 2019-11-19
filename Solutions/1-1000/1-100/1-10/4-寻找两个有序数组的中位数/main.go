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
