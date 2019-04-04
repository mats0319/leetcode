package main

import (
	"fmt"
)

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3) == "PAHNAPLSIIGYIR")
	fmt.Println(convert("PAYPALISHIRING", 4) == "PINALSIGYAHRPI")
}

func convert(s string, numRows int) string {
	var (
		sn     []int  = make([]int, len(s))
		result []byte = make([]byte, 0, len(sn))
	)

	var i int
	for i < len(sn) {
		for c := 1; c <= numRows && i < len(sn); c++ {
			sn[i] = c
			i++
		}
		for c := numRows - 1; c > 1 && i < len(sn); c-- {
			sn[i] = c
			i++
		}
	}
	for c := 1; c <= numRows; c++ {
		for i = 0; i < len(sn); i++ {
			if sn[i] == c {
				result = append(result, s[i])
			}
		}
	}
	return string(result)
}

// 
// Pick out each character in result string by calculate.
//
//func convert(s string, numRows int) string {
//	byts := make([]byte, 0, len(s))
//
//	for row := 0; row < numRows; row++ {
//		for i := 0; ; i++ {
//			k := zigzagIndexMap(numRows, row, i)
//			if k >= len(s) {
//				break
//			}
//
//			byts = append(byts, s[k])
//		}
//	}
//
//	return string(byts)
//}
//func zigzagIndexMap(n int, row int, i int) int {
//	if n == 1 {
//		return i
//	}
//
//	if row == 0 || row == n-1 {
//		return row + i*2*(n-1)
//	}
//
//	// row + ( (n-row-1) + (n-2-row+1) ) * ((i+1)/2) + (2*i)*(j/2)
//	return row + ((i+1)/2)*2*(n-row-1) + (i/2)*2*row
//}
