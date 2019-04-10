package main

import (
	"fmt"
)

func main() {
	fmt.Println(myAtoi("42") == int(42), myAtoi("42"), int(42))
	fmt.Println(myAtoi("   -42") == int(-42), myAtoi("   -42"), int(-42))
	fmt.Println(myAtoi("4193 with words") == int(4193), myAtoi("4193 with words"), int(4193))
	fmt.Println(myAtoi("words and 987") == int(0), myAtoi("words and 987"), int(0))
	fmt.Println(myAtoi("-91283472332") == int(-2147483648), myAtoi("-91283472332"), int(-2147483648))
	fmt.Println(myAtoi(" ") == int(0), myAtoi(" "), int(0))
	fmt.Println(myAtoi("   ") == int(0), myAtoi("   "), int(0))
	fmt.Println(myAtoi("9223372036854775808") == int(2147483647), myAtoi("9223372036854775808"), int(2147483647))
}

func myAtoi(str string) int {
	const (
		MAX = 2<<30 - 1
		MIN = -2 << 30
	)
	var (
		result int
		neg    = false
	)

	for {
		for i := range str {
			if str[i] != 32 {
				str = str[i:]
				break
			}
		}
		if len(str) == 0 {
			break
		}

		if str[0] == 43 {
		} else if str[0] == 45 {
			neg = true
		} else if str[0] >= 48 && str[0] <= 57 {
			result = int(str[0] - 48)
		} else {
			break
		}
		str = str[1:]

		for i := range str {
			if !(str[i] >= 48 && str[i] <= 57) {
				break
			}
			result = result*10 + int(str[i]-48)
			if result > MAX { // can not use >=
				break
			}
		}
		if neg {
			result = -result // int(0) == int(-0)
		}
		if result > MAX {
			result = MAX
		}
		if result < MIN {
			result = MIN
		}
		break
	}

	return result
}
