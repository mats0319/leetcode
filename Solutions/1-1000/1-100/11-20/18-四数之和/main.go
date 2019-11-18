package main

func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return nil
	}
	if len(nums) == 4 { // if len = 4, not sort, save time
		if sum(nums) == target {
			return [][]int{nums}
		} else {
			return nil
		}
	}

	sort(nums)

	var (
		length  = len(nums)
		a, b, c, d, t int
		result  = make([][]int, 0)
	)

	var temp1, temp2 int
	for ; a < length-3; a++ {
		temp1 = target - nums[a]

		for b = a + 1; b < length-2; b++ {
			temp2 = temp1 - nums[b]

			for c, d = b+1, length - 1;c < d; {
				t = temp2 - nums[c] - nums[d]
				if t == 0 {
					result = append(result, []int{nums[a], nums[b], nums[c], nums[d]})

					// skip repeated c and d
					for ; c < d && nums[c] == nums[c+1]; c++ {
					}
					for ; c < d && nums[d] == nums[d-1]; d-- {
					}

					c++
					d--
				} else if t < 0 {
					d--
				} else {
					c++
				}
			}

			// skip repeated b
			for b < length-2 && nums[b] == nums[b+1] {
				b++
			}
		}

        // skip repeated b
		for a < length-3 && nums[a] == nums[a+1] {
			a++
		}
	}

	return result
}

func sum(is []int) (sum int) {
	for i := range is {
		sum += is[i]
	}

	return
}

// quick sort
func sort(is []int) {
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

	sort(is[:small])
	sort(is[small+1:])

	return
}
