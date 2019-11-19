package main

func decodeString(s string) string {
	var (
		times, digits int
		strDecode     string
		indexStack    = make([]int, 0)
	)

	for i := 0; i < len(s); i++ {
		if s[i] == '[' {
			indexStack = append(indexStack, i)
		} else if s[i] == ']' {
			times, digits = getTimesAndDigits(s, indexStack[len(indexStack)-1])
			strDecode = strCopyTimes(s[indexStack[len(indexStack)-1]+1:i], times)

			s = s[:indexStack[len(indexStack)-1]-digits] + strDecode + s[i+1:] // modify s
			// old length: i - is[len(is)-1] + 1 + (times digits)
			// new length: len(strDecode)
			// offset: new - old = len(strDecode) - i + is[len(is)-1] - 1 - (times digits)
			// i += offset
			i += len(strDecode) + indexStack[len(indexStack)-1] - i - 1 - digits

			indexStack = indexStack[:len(indexStack)-1] // stack pop
		}
	}

	return s
}

// get valid int, before str[index] ('[')
func getTimesAndDigits(str string, index int) (times, digits int) {
	for j := index - 1; j >= 0 && str[j] >= '0' && str[j] <= '9'; j-- {
		digits++
	}

	for i := index - digits; i < index; i++ {
		times = times*10 + int(str[i]-'0')
	}

	return
}

func strCopyTimes(str string, times int) (result string) {
	for i := 0; i < times; i++ {
		result += str
	}

	return
}
