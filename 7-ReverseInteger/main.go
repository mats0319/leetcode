package main

import (
	"fmt"
)

func main() {
	fmt.Println(reverse(123) == int(321))
	fmt.Println(reverse(-123) == int(-321))
	fmt.Println(reverse(120) == int(21))
	fmt.Println(reverse(1534236469) == int(0))
}

func reverse(x int) int {
	const (
		MAX = 2<<30 - 1
		MIN = -2 << 30
	)
	fmt.Println(MAX, MIN)
	var result int

	for x != 0 {
		result = result*10 + x%10
		x /= 10
	}

	if result > MAX || result < MIN {
		result = 0
	}

	return result
}
