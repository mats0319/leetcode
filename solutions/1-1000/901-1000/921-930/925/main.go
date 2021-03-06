package mario

func isLongPressedName(name string, typed string) bool {
	if len(name) > len(typed) {
		return false
	}

	name = " " + name
	typed = " " + typed

	n, t := 1, 1
	isValid := true
	for n < len(name) && t < len(typed) {
		if name[n] != typed[t] && typed[t] != typed[t-1] {
			isValid = false
			break
		}

		if name[n] == typed[t] {
			n++
		}

		t++
	}

	for t < len(typed) && typed[t] == typed[t-1] {
		t++
	}

	return isValid && n >= len(name) && t >= len(typed)
}
