package mario

func countBinarySubstrings(s string) int {
	counts := make([]int, 0)
	{
		ct := 1
		for i := 0; i < len(s)-1; i++ {
			if s[i] == s[i+1] {
				ct++
			} else {
				counts = append(counts, ct)
				ct = 1
			}
		}

		counts = append(counts, ct)
	}

	var result int
	for i := 0; i < len(counts)-1; i++ {
		result += min(counts[i], counts[i+1])
	}

	return result
}

func min(a, b int) (result int) {
	if a < b {
		result = a
	} else {
		result = b
	}

	return result
}
