package mario

var m = map[byte]byte{
	'(': ')',
	'[': ']',
	'{': '}',
}

// s only contains 6 types of chars in m above
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
		if v, ok := m[s[i]]; ok { // left side char
			stack[top] = v
			top++
			continue
		}

		// right side char, not matched
		if top == 0 || s[i] != stack[top-1] {
			valid = false
			break
		}

		top--
	}

	return valid && top == 1
}
