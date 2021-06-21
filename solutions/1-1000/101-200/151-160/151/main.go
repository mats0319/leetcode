package mario

func reverseWords(s string) string {
	stack := make([]string, 0)
	s += " "
	for i, start := 0, -1; i < len(s); i++ {
		if s[i] != ' ' && start < 0 {
			start = i
		} else if s[i] == ' ' && start >= 0 {
			stack = append(stack, s[start:i])
			start = -1
		}
	}

	if len(stack) < 1 {
		return ""
	}

	res := ""
	for i := len(stack)-1; i >= 0; i-- {
		res += stack[i] + " "
	}

	return res[:len(res)-1]
}
