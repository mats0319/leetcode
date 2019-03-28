package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("cbbd"))
	fmt.Println(longestPalindrome("ac"))
}

func longestPalindrome(s string) string {
	length := len(s)
	if length == 0 || length == 1 {
		return s
	}

	var (
		resultStart  int
		resultLength int
		tempLength   int
	)
	for i := 0; i < length-1; i++ {
		if i > length-resultLength/2-1 {
			break
		}
		var (
			tempLengthSingleMiddle int = getLongestPalindromeLengthForAside(s, i, i)*2 + 1
			tempLengthDoubleMiddle int
		)
		if s[i] == s[i+1] {
			tempLengthDoubleMiddle = getLongestPalindromeLengthForAside(s, i, i+1)*2 + 2
		}
		if tempLengthSingleMiddle > tempLengthDoubleMiddle {
			tempLength = tempLengthSingleMiddle
		} else {
			tempLength = tempLengthDoubleMiddle
		}
		if tempLength <= 0 {
			tempLength = 1
		}

		if tempLength > resultLength {
			resultLength = tempLength
			resultStart = i - resultLength/2 - (resultLength%2 - 1)
		}
	}

	return s[resultStart : resultStart+resultLength]
}

func getLongestPalindromeLengthForAside(str string, i, j int) int {
	var c int
	for i >= 0 && j <= len(str)-1 && str[i] == str[j] {
		i--
		j++
		c++
	}
	return c - 1
}
