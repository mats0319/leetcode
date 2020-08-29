package mario

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	mulBytes := make([]byte, len(num1)+len(num2))

	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			add := (num1[i]-'0')*(num2[j]-'0') + mulBytes[i+j+1]

			mulBytes[i+j+1] = add % 10
			mulBytes[i+j] += add / 10
		}
	}

	for i := range mulBytes {
		mulBytes[i] += '0'
	}

	if mulBytes[0] == '0' {
		mulBytes = mulBytes[1:]
	}

	return string(mulBytes)
}
