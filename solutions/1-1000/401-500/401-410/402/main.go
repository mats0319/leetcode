package mario

// num.length >= k
func removeKdigits(num string, k int) string {
	if len(num) == k {
		return "0"
	}

	for k > 0 && len(num) > 0 {
		remove := false
		for i := 0; i < len(num)-1; i++ {
			if num[i] > num[i+1] {
				num = num[:i] + num[i+1:]
				if i == 0 {
					j := 0
					for j < len(num) && num[j] == '0' {
						j++
					}
					num = num[j:]
				}

				remove = true
				break
			}
		}

		if !remove {
			num = num[:len(num)-1]
		}

		k--
	}

	if len(num) == 0 {
		num = "0"
	}

	return num
}
