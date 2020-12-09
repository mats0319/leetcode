package mario

func sortString(s string) string {
	counts := make([]uint8, 26)
	for _, char := range s {
		counts[int(char-'a')]++
	}

	sum := len(s)
	res := ""
	for sum > 0 {
		for i := 0; i < len(counts); i++ {
			if counts[i] > 0 {
				counts[i]--
				sum--
				res += string(byte(i) + 'a')
			}
		}

		for i := len(counts) - 1; i >= 0; i-- {
			if counts[i] > 0 {
				counts[i]--
				sum--
				res += string(byte(i) + 'a')
			}
		}
	}

	return res
}
