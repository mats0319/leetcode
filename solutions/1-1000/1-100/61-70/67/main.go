package mario

// no leading zero
func addBinary(a string, b string) string {
	length := max(len(a), len(b)) + 1
	res := make([]byte, length)

	ia := len(a) - 1
	ib := len(b) - 1

	carryDigit := 0
	for i := length - 1; i >= 0; i-- {
		count := carryDigit

		if ia >= 0 && a[ia] == '1' {
			count++
		}
		if ib >= 0 && b[ib] == '1' {
			count++
		}

		switch count {
		case 0:
			res[i] = '0'
			carryDigit = 0
		case 1:
			res[i] = '1'
			carryDigit = 0
		case 2:
			res[i] = '0'
			carryDigit = 1
		case 3:
			res[i] = '1'
			carryDigit = 1
		}

		ia--
		ib--
	}

	if res[0] == '0' {
		res = res[1:]
	}

	return string(res)
}

func max(a, b int) (res int) {
	if a > b {
		res = a
	} else {
		res = b
	}

	return
}
