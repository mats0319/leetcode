package main

import (
	"fmt"
)

func main() {
	fmt.Println(romanToInt("III") == 3)
	fmt.Println(romanToInt("IV") == 4)
	fmt.Println(romanToInt("IX") == 9)
	fmt.Println(romanToInt("LVIII") == 58)
	fmt.Println(romanToInt("MCMXCIV") == 1994)
	fmt.Println(romanToInt("XI") == 11)
}

func romanToInt(s string) int {
	var (
		convertValue  = []int{1, 5, 10, 50, 100, 500, 1000}
		convertSToInt = make([]int, len(s))
		result        int
	)

	for i := range s {
		switch s[i] {
		case 'I':
			convertSToInt[i] = 0
		case 'V':
			convertSToInt[i] = 1
		case 'X':
			convertSToInt[i] = 2
		case 'L':
			convertSToInt[i] = 3
		case 'C':
			convertSToInt[i] = 4
		case 'D':
			convertSToInt[i] = 5
		case 'M':
			convertSToInt[i] = 6
		}
	}

	for i := 0; i < len(convertSToInt); i++ {
		if i < len(convertSToInt)-1 && convertSToInt[i] < convertSToInt[i+1] {
			result += convertValue[convertSToInt[i+1]] - convertValue[convertSToInt[i]]
			i++
		} else {
			result += convertValue[convertSToInt[i]]
		}
	}

	return result
}
