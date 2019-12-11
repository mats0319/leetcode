package mario

var m = map[byte]byte{
	'(': ')',
	'[': ']',
	'{': '}',
}

func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}

	var (
		stack = make([]byte, len(s)+1)
		top   = 1 // skip stack[0], for keep available when calculate 'stack[top-1]' below
		valid = true
	)

	for i := range s {
		if isLeft(s[i]) {
			stack[top] = m[s[i]]
			top++
		} else if isRight(s[i]) && stack[top-1] == s[i] {
			if top == 0 {
				valid = false
				break
			}
			top--
		}
	}

	return valid && top == 1
}

func isLeft(b byte) bool {
    _, ok := m[b]
    return ok
}

func isRight(b byte) bool {
	return b == ')' || b == ']' || b == '}'
}
