package mario

import (
	"strings"
)

func reverseWords(s string) string {
	words := strings.Split(s, " ")

	for i := range words {
		convertBytes := make([]byte, 0, len(words[i]))
		for j := len(words[i])-1; j >= 0; j-- {
			convertBytes = append(convertBytes, words[i][j])
		}
		words[i] = string(convertBytes)
	}

	return strings.Join(words, " ")
}
