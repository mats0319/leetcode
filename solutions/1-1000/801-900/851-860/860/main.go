package mario

func lemonadeChange(bills []int) bool {
	count := make([]int, 2)

	isValid := true
ALL:
	for _, v := range bills {
		switch v {
		case 5:
			count[0]++
		case 10:
			count[1]++
			count[0]--
			if count[0] < 0 {
				isValid = false
				break ALL
			}
		case 20:
			if count[1] > 0 && count[0] > 0 {
				count[1]--
				count[0]--
			} else if count[0] >= 3 {
				count[0] -= 3
			} else {
				isValid = false
				break ALL
			}
		}
	}

	return isValid
}
