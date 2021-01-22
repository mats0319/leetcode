package mario

func isPalindrome(s string) bool {
	isValid := true
	for left, right := 0, len(s)-1; left < right; left, right = left+1, right-1 {
		for left < len(s) && !isValidChar(s[left]) {
			left++
		}
		for right >= 0 && !isValidChar(s[right]) {
			right--
		}

		if left > right {
			break
		}

		if toLowerCase(s[left]) != toLowerCase(s[right]) {
			isValid = false
			break
		}
	}

	return isValid
}

func isValidChar(char byte) bool {
	return ('A' <= char && char <= 'Z') || ('a' <= char && char <= 'z') || ('0' <= char && char <= '9')
}

func toLowerCase(char byte) byte {
	res := char
	if 'A' <= char && char <= 'Z' {
		res += 'a'-'A'
	}

	return res
}
