package mario

import "strings"

func wordPattern(pattern string, s string) bool {
	strs := strings.Split(s, " ")
	if len(pattern) != len(strs) {
		return false
	}

	byteMap := make(map[byte]string, len(pattern))
	strExist := make(map[string]bool, len(pattern))
	isMatched := true
	for i := 0; i < len(pattern); i++ {
		v, ok := byteMap[pattern[i]]

		if (ok && strs[i] != v) || (!ok && strExist[strs[i]]) {
			isMatched = false
			break
		} else if !ok {
			byteMap[pattern[i]] = strs[i]
			strExist[strs[i]] = true
		}
	}

	return isMatched
}
