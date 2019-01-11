package main

import (
	"fmt"
)

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}

func lengthOfLongestSubstring(s string) int {
	slen := len(s)
	var (
		m              map[uint8]int = make(map[uint8]int)
		length, result int
		ok             bool
	)
	for i := 0; i < slen; i++ {
		if _, ok = m[s[i]]; !ok {
			m[s[i]] = 0
			length++
		} else {
			if i -= length; i >= slen-result-1 {
				break
			}
			length = 0
			m = make(map[uint8]int)
		}
		if result < length {
			result = length
		}
	}
	return result
}
