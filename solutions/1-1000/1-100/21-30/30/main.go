package mario

func findSubstring(s string, words []string) []int {
	if len(words) < 1 || len(s) < len(words)*len(words[0]) {
		return nil
	}

	length := len(words) * len(words[0])
	res := make([]int, 0, len(s)-length+1)
	for i := 0; i < len(s)-length+1; i++ {
		if matched(s[i:i+length], words) {
			res = append(res, i)
		}
	}

	return res
}

func matched(str string, words []string) bool {
	left := words
	isMatched := true
	for i := 0; i < len(str); i += len(words[0]) {
		index := getIndex(str[i:i+len(words[0])], left)
		if index < 0 {
			isMatched = false
			break
		}

		left[index], left[len(left)-1] = left[len(left)-1], left[index]
		left = left[:len(left)-1]
	}

	return isMatched
}

func getIndex(str string, array []string) int {
	index := -1
	for i, v := range array {
		if str == v {
			index = i
			break
		}
	}

	return index
}
