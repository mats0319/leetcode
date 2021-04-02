package mario

import "strconv"

// input is always valid
func calculate(s string) int {
	return calcSimpleExpression(trimAllSpace(s))
}

func calcSimpleExpression(str string) (res int) {
	var operator byte
	for i := 0; i < len(str); i++ {
		switch str[i] {
		case '(':
			j := i + 1
			leftParentheses := 1
			for j < len(str) && leftParentheses > 0 {
				if str[j] == '(' {
					leftParentheses++
				} else if str[j] == ')' {
					leftParentheses--
				}

				j++
			}

			value := calcSimpleExpression(str[i+1 : j-1])
			switch operator {
			case 0:
				res = value
			case '+':
				res += value
			case '-':
				res -= value
			}

			i = j - 1
		case '+', '-':
			operator = str[i]
		default: // numbers
			j := i + 1
			for j < len(str) && '0' <= str[j] && str[j] <= '9' {
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

	return
}

func trimAllSpace(s string) string {
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			s = s[:i] + s[i+1:]
			i--
		}
	}

	return s
}
