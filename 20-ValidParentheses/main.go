package main

import (
	"fmt"
)

func main() {
	fmt.Println(isValid("()") == true)
	fmt.Println(isValid("()[]{}") == true)
	fmt.Println(isValid("") == true)
	fmt.Println(isValid("(]") == false)
	fmt.Println(isValid("([)]") == false)
	fmt.Println(isValid("{[]}") == true)
	fmt.Println(isValid("}{") == false)
	fmt.Println(isValid("((") == false)
}

func isValid(s string) bool {
	length := len(s)
	if length%2 == 1 {
		return false
	}

	stack := make([]byte, length+1)
	length = 0

	for i := range s {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack[length] = s[i]
			length++
		} else {
			if length <= 0 {
				return false
			}

			if s[i] == stack[length-1]+stack[length-1]%2+1 {
				length--
			} else {
				return false
			}
		}
	}

	return length == 0
}
