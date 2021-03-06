package mario

func countAndSay(n int) string {
	str := "1"
	for i := 2; i <= n; i++ {
		str = describe(str)
	}

	return str
}

func describe(str string) string {
	char := str[0]
	var count byte = 1

	res := ""
	for i := 1; i < len(str); i++ {
		if str[i] == char {
			count++
			continue
		}

		res += string(count+'0') + string(char)
		char = str[i]
		count = 1
	}

	return res + string(count+'0') + string(char)
}
