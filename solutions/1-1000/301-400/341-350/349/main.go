package mario

import "github.com/mats9693/utils/sort"

func intersection(nums1 []int, nums2 []int) []int {
	if len(nums1) < 1 || len(nums2) < 1 {
		return nil
	}

	sort.QuickSort(nums1)
	sort.QuickSort(nums2)

	result := make([]int, 0)
	pre := nums1[0] - 1
	for m, n := 0, 0; m < len(nums1) && n < len(nums2); {
		if nums1[m] == nums2[n] {
			if nums1[m] != pre {
				result = append(result, nums1[m])
				pre = nums1[m]
			}

			m++
			n++
		} else if nums1[m] < nums2[n] {
			m++
		} else if nums1[m] > nums2[n] {
			n++
		}
	}

	return result
}
