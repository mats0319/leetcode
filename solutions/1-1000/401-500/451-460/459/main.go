package mario

func repeatedSubstringPattern(s string) (result bool) {
	for i := 1; i <= len(s)/2; i++ {
		if len(s)%i == 0 {
			frequency := len(s) / i
			var str string
			for j := 0; j < frequency; j++ {
				str += s[:i]
			}

			if str == s {
				result = true
				break
			}
		}
	}

	return
}
