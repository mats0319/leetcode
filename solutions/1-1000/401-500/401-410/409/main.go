package mario

func longestPalindrome(s string) int {
	charMap := make(map[byte]int, 0)
	for i := range s {
		charMap[s[i]]++
	}

	res := 0
	validSingleCenter := 0
	for i := range charMap {
		times := charMap[i]
		if times%2 == 1 {
			validSingleCenter = 1
		}

		if times >= 2 {
			res += times - times%2
		}
	}

	return res + validSingleCenter
}
