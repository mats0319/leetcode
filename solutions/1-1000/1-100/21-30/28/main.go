package mario

func strStr(haystack string, needle string) int {
	if len(needle) < 1 {
		return 0
	}

	index := -1
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		if haystack[i] == needle[0] {
			if isMatched(haystack[i:i+len(needle)], needle) {
				index = i
				break
			}
		}
	}

	return index
}

// a.length == b.length
func isMatched(a, b string) bool {
	isEqual := true
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			isEqual = false
			break
		}
	}

	return isEqual
}
