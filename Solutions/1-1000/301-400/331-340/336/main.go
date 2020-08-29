package mario

func palindromePairs(words []string) [][]int {
	if len(words) <= 1 {
		return nil
	}

	result := make([][]int, 0)
	for i := 0; i < len(words)-1; i++ {
		for j := i + 1; j < len(words); j++ {
			if isPalindromePair(words[i], words[j]) {
				result = append(result, []int{i, j})
			}
			if isPalindromePair(words[j], words[i]) {
				result = append(result, []int{j, i})
			}
		}
	}

	return result
}

func isPalindromePair(str1, str2 string) bool {
	if str1 == "" && str2 == "" {
		return true
	} else if str1 == "" {
		return isPalindrome(str2)
	} else if str2 == "" {
		return isPalindrome(str1)
	}

	start, end := 0, len(str2)-1
	isPalindromeFlag := true
	for ; start < len(str1) && end >= 0; start, end = start+1, end-1 {
		if str1[start] != str2[end] {
			isPalindromeFlag = false
			break
		}
	}

	if isPalindromeFlag {
		if len(str1) > len(str2) {
			isPalindromeFlag = isPalindrome(str1[start:])
		} else if len(str1) < len(str2) {
			isPalindromeFlag = isPalindrome(str2[:end+1])
		}
	}

	return isPalindromeFlag
}

func isPalindrome(str string) bool {
	if len(str) <= 1 { // str.length == 2, str[0] == str[1] is included below, or even str.length == 1 is.
		return true
	}

	isPalindromeFlag := true
	for start, end := 0, len(str)-1; start <= end; start, end = start+1, end-1 {
		if str[start] != str[end] {
			isPalindromeFlag = false
			break
		}
	}

	return isPalindromeFlag
}
