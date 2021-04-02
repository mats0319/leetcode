package mario

func longestConsecutive(nums []int) int {
	m := make(map[int]struct{}, len(nums))
	for i := range nums {
		m[nums[i]] = struct{}{}
	}

	res := 0
	for len(m) > 0 {
	    var k int
	    for k = range m {
	    	delete(m, k)
	        break
        }

		length := 1

		for i := k - 1; ; i-- {
			if _, ok := m[i]; !ok {
				break
			}

			length++
			delete(m, i)
		}

		for i := k + 1; ; i++ {
			if _, ok := m[i]; !ok {
				break
			}

			length++
			delete(m, i)
		}

		if res < length {
			res = length
		}
	}

	return res
}
