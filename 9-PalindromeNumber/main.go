package main

import (
	"fmt"
)

func main() {
	fmt.Println(isPalindrome(121) == true)
	fmt.Println(isPalindrome(-121) == false)
	fmt.Println(isPalindrome(10) == false)
}

func isPalindrome(x int) bool {
	var (
		reverse      int
		isPalindrome bool
	)
	for {
		if x < 0 || (x%10 == 0 && x != 0) {
			break
		}
		for x > reverse {
			reverse = reverse*10 + x%10
			x /= 10
		}
		if x == reverse || x == reverse/10 {
			isPalindrome = true
		}
		break
	}

	return isPalindrome
}
