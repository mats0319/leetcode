package mario

import "strconv"

func calculate2(s string) int {
	str := trimAllMeaninglessChars(s)
	res := 0
	var operator byte
	for i := 0; i < len(str); i++ {
		switch str[i] {
		case '+', '-':
			operator = str[i]
		default:
			j := i + 1
			for j < len(str) && isNum(str[j]) {
				j++
			}
			value, _ := strconv.Atoi(str[i:j])
			switch operator {
			case 0:
				res = value
			case '+':
				res += value
			case '-':
				res -= value
			}
			i = j - 1
		}
	}

	return res
}

func trimAllMeaninglessChars(s string) string {
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' || s[i] == '(' || s[i] == ')' {
			s = s[:i] + s[i+1:]
			i--
		}
	}

	return s
}
