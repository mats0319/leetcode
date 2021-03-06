package mario

func reverseWords(s string) string {
	stack := make([]string, 0)
	start := 0
	write := false

	s += " "
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case ' ':
			if !write {
				break
			}

			stack = append(stack, " "+s[start:i])
			start = i
			write = false
		default:
			if !write {
				start = i
				write = true
			}
		}
	}

	if len(stack) < 1 {
		return ""
	}

	res := ""
	for i := len(stack) - 1; i >= 0; i-- {
		res += stack[i]
	}

	return res[1:]
}
