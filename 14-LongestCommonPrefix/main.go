package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}) == "fl")
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}) == "")
	fmt.Println(longestCommonPrefix([]string{}) == "")
	fmt.Println(longestCommonPrefix([]string{"c", "c"}) == "c")
	fmt.Println(longestCommonPrefix([]string{"aaa", "aa", "aaa"}) == "aa")
}

func longestCommonPrefix(strs []string) string {
	var (
		shortestStrIndex int
		i                int
		result           string
	)
	for {
		if len(strs) == 0 || hasEmptyStr(strs, &shortestStrIndex) || strs == nil {
			break
		}
		if len(strs) == 1 {
			result = strs[0]
			break
		}
		var notMatch bool
		for i = range strs[shortestStrIndex] {
			for is := range strs {
				if strs[is][i] != strs[shortestStrIndex][i] {
					notMatch = true
					break
				}
			}
			if notMatch {
				break
			}
		}
		if !notMatch {
			i++
		}
		result = strs[shortestStrIndex][:i]
		break
	}

	return result
}

func hasEmptyStr(strs []string, ssi *int) bool {
	var emptyStr bool
	for i := range strs {
		if strs[i] == "" {
			emptyStr = true
			break
		}
		if len(strs[i]) < len(strs[*ssi]) {
			*ssi = i
		}
	}
	return emptyStr
}
