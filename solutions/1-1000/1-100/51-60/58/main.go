package mario

func lengthOfLastWord(s string) int {
	length := 0
	inWord := false
	for i := len(s) - 1; i >= 0; i-- {
		if isLetter(s[i]) {
			inWord = true
			length++
		} else if inWord {
			break
		}
	}

	return length
}

func isLetter(c byte) bool {
	return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z')
}
