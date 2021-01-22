package mario

// nums.length >= 3
func maximumProduct(nums []int) int {
	QuickSort(nums)

	return max(nums[len(nums)-1]*nums[0]*nums[1], nums[len(nums)-1]*nums[len(nums)-2]*nums[len(nums)-3])
}

func max(a, b int) (res int) {
	if a > b {
		res = a
	} else {
		res = b
	}

	return
}

// QuickSort quick sort
func QuickSort(is []int) {
	if len(is) <= 1 {
		return
	}

	var small int
	{
		var big int
		for i := 1; i < len(is); i++ {
			if is[i] > is[0] {
				big++
			} else {
				small++
				if big != 0 {
					is[i], is[small] = is[small], is[i]
				}
			}
		}
		if small != 0 {
			is[0], is[small] = is[small], is[0]
		}
	}

	QuickSort(is[:small])
	QuickSort(is[small+1:])

	return
}
