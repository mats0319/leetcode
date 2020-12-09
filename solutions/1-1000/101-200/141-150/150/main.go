package mario

import "strconv"

func evalRPN(tokens []string) int {
	stack := make([]int, 0, len(tokens))
	for _, str := range tokens {
		switch str {
		case "+", "-", "*", "/":
			length := len(stack)
			stack = append(stack[:length-2], calc(stack[length-2], stack[length-1], str))
		default: // number
			v, _ := strconv.Atoi(str)
			stack = append(stack, v)
		}
	}

	return stack[0]
}

func calc(operandOne, operandTwo int, method string) int {
	switch method {
	case "+":
		operandOne += operandTwo
	case "-":
		operandOne -= operandTwo
	case "*":
		operandOne *= operandTwo
	case "/":
		operandOne /= operandTwo
	}

	return operandOne
}
