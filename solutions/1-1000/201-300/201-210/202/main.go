package mario

func isHappy(n int) bool {
	value := []int{n}
	isHappyFlag := false
	for {
		sum := 0
		for n != 0 {
			digit := n % 10
			sum += digit * digit
			n /= 10
		}

		if sum == 1 {
			isHappyFlag = true
			break
		}

		if isExist(value, sum) {
			break
		}

		value = append(value, sum)
		n = sum
	}

	return isHappyFlag
}

func isExist(array []int, value int) bool {
	isExistFlag := false
	for i := range array {
		if array[i] == value {
			isExistFlag = true
			break
		}
	}

	return isExistFlag
}
