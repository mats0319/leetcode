package mario

func addStrings(num1 string, num2 string) string {
	if num1 == "" {
		return num2
	} else if num2 == "" {
		return num1
	}

	// make sure num1 is more short
	if len(num1) > len(num2) {
		num1, num2 = num2, num1
	}

	var (
		i, j, k = len(num1) - 1, len(num2) - 1, len(num2) - 1
		t, flag int
		result  = make([]byte, len(num2), len(num2))
	)
	for ; i >= 0; i, j, k = i-1, j-1, k-1 {
		t = int(num1[i]-'0') + int(num2[j]-'0') + flag

		result[k] = byte(t%10 + '0')
		flag = t / 10
	}

	for ; j >= 0 && flag == 1; j, k = j-1, k-1 {
		t = int(num2[j]-'0') + flag

		result[k] = byte(t%10 + '0')
		flag = t / 10
	}

	if flag == 1 {
		return "1" + string(result)
	} else if j != -1 {
		return num2[:j+1] + string(result[j+1:])
	}

	return string(result)
}
