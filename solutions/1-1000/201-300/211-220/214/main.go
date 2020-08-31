package mario

func shortestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}

	str := s
	if len(str)%2 == 1 {
		str += " "
	}

	var paddingStart int
	for i := (len(str)+1)/2 - 1; i >= 0; i-- {
		if str[i] == str[i+1] && isPalindrome(str[:i*2+2]) {
			paddingStart = i*2 + 2
			break
		}

		if isPalindrome(str[:i*2+1]) {
			paddingStart = i*2 + 1
			break
		}
	}

	prefixBytes := make([]byte, 0)
	for i := len(s) - 1; i >= paddingStart; i-- {
		prefixBytes = append(prefixBytes, s[i])
	}

	return string(prefixBytes) + s
}

func isPalindrome(str string) bool {
	// caller make sure str will not be empty

	l, r := 0, len(str)-1
	for ; l < r && str[l] == str[r]; l, r = l+1, r-1 {
	}

	return l >= r
}
