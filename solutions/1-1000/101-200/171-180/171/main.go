package mario

func titleToNumber(columnTitle string) int {
	sum := 0
	digit := 1
	for i := len(columnTitle) - 1; i >= 0; i-- {
		char := int(columnTitle[i] - 'A' + 1)

		sum += char * digit
		digit *= 26
	}

	return sum
}
