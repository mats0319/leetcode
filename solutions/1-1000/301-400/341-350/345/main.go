package mario

func reverseVowels(s string) string {
	byteSlice := []byte(s)

	for left, right := 0, len(byteSlice)-1; ; left, right = left+1, right-1 {
		for left < right && !isValid(byteSlice[left]) {
			left++
		}

		for left < right && !isValid(byteSlice[right]) {
			right--
		}

		if left >= right {
			break
		}

		byteSlice[left], byteSlice[right] = byteSlice[right], byteSlice[left]
	}

	return string(byteSlice)
}

func isValid(b byte) bool {
	return b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u' ||
		b == 'A' || b == 'E' || b == 'I' || b == 'O' || b == 'U'
}
