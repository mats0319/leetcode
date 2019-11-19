package main

var matchMap = map[byte][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

func letterCombinations(digits string) (result []string) {
	if digits == "" {
		return
	}

	var t []string

	result = matchMap[digits[0]]
	for i := 0; i < len(digits)-1; i++ {
		t = nil
		for j := range result {
			for k := range matchMap[digits[i+1]] {
				t = append(t, result[j]+matchMap[digits[i+1]][k])
			}
		}
		result = t
	}

	return
}
