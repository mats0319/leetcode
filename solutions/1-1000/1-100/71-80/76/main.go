package mario

// both s and t not nil
func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}

	tMap := convertToMap(t)
    sum := len(t)

	left, right := -1, 0
	for ; right < len(s) && sum > 0; right++ {
        char := s[right]

        _, ok := tMap[char]
		if !ok {
            continue
        }

        tMap[char]--
        sum--
        if left < 0 { // only run once, init left
            left = right
        }
	}

    if sum > 0 {
        return ""
    }

    res := s[left: right]

    // todo: right counter go through

	return res
}

func convertToMap(str string) map[byte]int {
	res := make(map[byte]int)
	for i := range str {
		res[str[i]]++
	}

	return res
}
