package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbba"))
}

// 20ms,6.1M memory submission,which use the least memory in Go's submission.
func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}

	rstr := ""
	cstr := fmt.Sprintf("%c", s[0])
	for _, i := range s[1:] {
		k := strings.LastIndex(cstr, string(i))
		if k == -1 {
			cstr += fmt.Sprintf("%s", string(i))
		} else {
			if len(cstr) > len(rstr) {
				rstr = cstr
			}
			cstr = cstr[k+1:] + string(i)
		}
	}

	if len(rstr) >= len(cstr) {
		return len(rstr)
	} else {
		return len(cstr)
	}
}

// 4ms submission.
//func lengthOfLongestSubstring(s string) int {
//	var index_array [256]int
//	for i:=0;i<256;i++{
//		index_array[i]=-1
//	}
//
//	var head int
//	max_length := 0
//	for pos, char := range s{
//		if index_array[char] == -1 {
//			index_array[char] = pos
//		}else{
//			head = max(index_array[char] + 1, head)
//			index_array[char] = pos
//		}
//		max_length = max(max_length, pos-head+1)
//	}
//	return max_length
//}
//
//func max(a, b int) int{
//	if a > b {
//		return a
//	}
//	return b
//}

// my submission...
//func lengthOfLongestSubstring(s string) int {
//	slen := len(s)
//	var (
//		m              map[uint8]int = make(map[uint8]int)
//		length, result int
//		ok             bool
//	)
//	for i := 0; i < slen; i++ {
//		if _, ok = m[s[i]]; !ok {
//			m[s[i]] = 0
//			length++
//		} else {
//			if i -= length; i >= slen-result-1 {
//				break
//			}
//			length = 0
//			m = make(map[uint8]int)
//		}
//		if result < length {
//			result = length
//		}
//	}
//	return result
//}
