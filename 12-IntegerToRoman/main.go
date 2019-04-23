package main

import "fmt"

func main() {
	fmt.Println(intToRoman(3) == "III")
	fmt.Println(intToRoman(4) == "IV")
	fmt.Println(intToRoman(9) == "IX")
	fmt.Println(intToRoman(58) == "LVIII")
	fmt.Println(intToRoman(1994) == "MCMXCIV")
	fmt.Println(intToRoman(11) == "XI")
}

func intToRoman(num int) string {
	var (
		// value:                1    5   10   50   100  500  1000
		convertArray = []string{"I", "V", "X", "L", "C", "D", "M"}
		// index:                0    1    2    3    4    5    6
		result string
		digit  int
		temp   string
	)
	for i := 0; num != 0; i++ {
		digit = num % 10
		num /= 10
		if digit == 4 {
			result = convertArray[i*2] + convertArray[i*2+1] + result
			continue
		}
		if digit == 9 {
			result = convertArray[i*2] + convertArray[i*2+2] + result
			continue
		}
		temp = ""
		if digit > 4 {
			temp = convertArray[i*2+1]
			digit -= 5
		}
		for ; digit > 0; digit-- {
			temp += convertArray[i*2]
		}
		result = temp + result
	}
	return result
}

// enumeration: (least memory)
//func intToRoman(num int) string {
//	M := []string{"", "M", "MM", "MMM"}
//	C := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
//	X := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
//	I := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
//	return M[num/1000] + C[num%1000/100] + X[num%100/10] + I[num%10]
//}
