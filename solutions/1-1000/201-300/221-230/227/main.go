package mario

import "strconv"

func calculate(s string) int {
	s = trimAllSpace(s)

	number := make([]int, 0, len(s))
	operator := make([]byte, 0, len(s))

	calcFlag := false
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '+', '-':
			operator = append(operator, s[i])
		case '*', '/':
			operator = append(operator, s[i])
			calcFlag = true
		default: // number
			j := i+1
			for j < len(s) && isNum(s[j]) {
				j++
			}

			v, _ := strconv.Atoi(s[i:j])
			i = j-1

			if !calcFlag {
				number = append(number, v)
				break
			}

			number = append(number[:len(number)-1], calc(number[len(number)-1], v, operator[len(operator)-1]))
			operator = operator[:len(operator)-1]

			calcFlag = false
		}
	}

	operandOne := number[0]
	opcode := 0
	for i := 1; i < len(number); i++ {
		operandOne = calc(operandOne, number[i], operator[opcode])
		opcode++
	}

	return operandOne
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

func calc(operandOne, operandTwo int, operator byte) int {
	switch operator {
	case '+':
		operandOne += operandTwo
	case '-':
		operandOne -= operandTwo
	case '*':
		operandOne *= operandTwo
	case '/':
		operandOne /= operandTwo
	}

	return operandOne
}

func isNum(n byte) bool {
	return '0' <= n && n <= '9'
}
