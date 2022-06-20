package mario

// find substr in s, which length is 10 and appear more than once
func findRepeatedDnaSequences(s string) []string {
	if len(s) < 10 {
		return nil
	}

	m := make(map[string]int, 0)
	for i := 10; i <= len(s); i++ {
		m[s[i-10:i]]++
	}

	res := make([]string, 0, len(m))
	for str, times := range m {
		if times > 1 {
			res = append(res, str)
		}
	}

	return res
}
