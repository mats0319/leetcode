package mario

// only contains lower-case letters
func firstUniqChar(s string) int {
	index := [26]int{}
	// init
	for i := 0; i < len(index); i++ {
		index[i] = -1
	}

	for i := 0; i < len(s); i++ {
		if index[s[i]-'a'] == -1 {
			index[s[i]-'a'] = i
		} else {
			index[s[i]-'a'] = -2 // may re-set '-2' lots of times, but ok
		}
	}

	res := -1
	for i := 0; i < len(index); i++ {
		if index[i] >= 0 { // appear only once
			if res < 0 {
				res = index[i]
			} else if index[i] < res {
				res = index[i]
			}
		}
	}

	return res
}
