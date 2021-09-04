package mario

func canConstruct(ransomNote string, magazine string) bool {
	have := [26]int{}
	for _, r := range magazine {
		have[r-'a']++
	}

	isValid := true
	for _, r := range ransomNote {
		index := r - 'a'

		have[index]--
		if have[index] < 0 {
			isValid = false
			break
		}
	}

	return isValid
}
