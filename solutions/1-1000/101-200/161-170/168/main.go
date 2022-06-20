package mario

var letters []string = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
	"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

func convertToTitle(columnNumber int) string {
	res := ""
	for {
		if 1 <= columnNumber && columnNumber <= 26 {
			res = letters[columnNumber-1] + res
			break
		}

		n := columnNumber / 26
		mod := columnNumber % 26
		if mod == 0 {
			n -= 1
			mod = 26
		}

		res = letters[mod-1] + res

		columnNumber = n
	}

	return res
}
