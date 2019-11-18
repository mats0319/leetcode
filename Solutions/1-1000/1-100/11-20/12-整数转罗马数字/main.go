package main

func intToRoman(num int) string {
	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	ints := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	var result string
	for i := 0; i < 13; i++ {
		for num >= ints[i] {
			num -= ints[i]
			result += romans[i]
		}
	}

	return result
}
