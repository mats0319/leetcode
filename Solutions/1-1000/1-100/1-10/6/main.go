package mario

import "strings"

// optimize: decrease run time
func convert(s string, numRows int) string {
	if numRows == 1 || len(s) <= numRows {
		return s
	}

	var (
		strSlice = make([]string, numRows)
		currRow  int
		offset   = -1
	)
	for i := 0; i < len(s); i++ {
		strSlice[currRow] += string(s[i])
		if currRow == 0 || currRow == numRows-1 {
			offset = -offset
		}
		currRow += offset
	}

	return strings.Join(strSlice, "")
}

// Pick out each character in result string by calculate.
func convert2(s string, numRows int) string {
	byts := make([]byte, 0, len(s))

	for row := 0; row < numRows; row++ {
		for i := 0; ; i++ {
			k := zigzagIndexMap(numRows, row, i)
			if k >= len(s) {
				break
			}

			byts = append(byts, s[k])
		}
	}

	return string(byts)
}

func zigzagIndexMap(n int, row int, i int) int {
	if n == 1 {
		return i
	}

	if row == 0 || row == n-1 {
		return row + i*2*(n-1)
	}

	// row + ( (n-row-1) + (n-2-row+1) ) * ((i+1)/2) + (2*i)*(j/2)
	return row + ((i+1)/2)*2*(n-row-1) + (i/2)*2*row
}
