package mario

func removeDuplicateLetters(s string) string {
	count := [26]int{}
	for _, v := range s {
		count[int(v-'a')]++
	}

	exist := [26]bool{}
	stack := make([]byte, 0, 26)
	for i := 0; i < len(s); i++ {
		count[int(s[i]-'a')]--

		if exist[int(s[i]-'a')] {
			continue
		}

		for len(stack) > 0 && stack[len(stack)-1] > s[i] && count[int(stack[len(stack)-1]-'a')] > 0 {
			exist[int(stack[len(stack)-1]-'a')] = false
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, s[i])
		exist[int(s[i]-'a')] = true
	}

	return string(stack)
}
