package mario

func addStrings(num1 string, num2 string) string {
	if num1 == "" {
		return num2
	} else if num2 == "" {
		return num1
	}

	if len(num1) > len(num2) { // make sure num1 is more short
		num1, num2 = num2, num1
	}

	var (
		i, j, k        = len(num1) - 1, len(num2) - 1, len(num2) - 1 // k is index for 'arb'
		addResultBytes = make([]byte, len(num2), len(num2))
		addCarryFlag   uint8
	)
	for ; i >= 0; i, j, k = i-1, j-1, k-1 {
		t := (num1[i] - '0') + (num2[j] - '0') + addCarryFlag

		addCarryFlag = t / 10
		addResultBytes[k] = t%10 + '0'
	}

	for ; j >= 0 && addCarryFlag == 1; j, k = j-1, k-1 {
		t := (num2[j] - '0') + addCarryFlag

		addCarryFlag = t / 10
		addResultBytes[k] = t%10 + '0'
	}

	var result string
	if addCarryFlag == 1 {
		result = "1" + string(addResultBytes)
	} else if j >= 0 {
		result = num2[:j+1] + string(addResultBytes[j+1:])
	} else {
		result = string(addResultBytes)
	}

	return result
}
