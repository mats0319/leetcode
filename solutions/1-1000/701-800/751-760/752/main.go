package mario

var operateLoop = []byte("0123456789")

func openLock(deadends []string, target string) int {
	if target == "0000" {
		return 0
	}
	if contains(deadends, "0000") {
		return -1
	}

	visitedMap := make(map[string]struct{})

	options := []string{"0000"}
	operateTimes := -1

ALL:
	for times := 1; len(options) > 0; times++ {
		nextOptions := make([]string, 0)

		for len(options) > 0 {
			item := options[0]
			options = options[1:]

			for _, v := range operateOnce(item) {
				if v == target {
                    operateTimes = times
					break ALL
				}

				if _, ok := visitedMap[v]; !ok && !contains(deadends, v) {
					visitedMap[v] = struct{}{}
					nextOptions = append(nextOptions, v)
				}
			}
		}

		options = nextOptions
	}

	return operateTimes
}

func operateOnce(item string) []string {
	res := make([]string, 0, len(item)*2)
	for i := 0; i < len(item); i++ {
		for _, v := range operateOneDigit(item[i]) {
			res = append(res, item[:i]+v+item[i+1:])
		}
	}

	return res
}

func operateOneDigit(char byte) []string {
	length := len(operateLoop)

	res := make([]string, 0, 2)
	for i := 0; i < length; i++ {
		if operateLoop[i] == char {
			res = append(res, string(operateLoop[(i+1+length)%length]), string(operateLoop[(i-1+length)%length]))
			break
		}
	}

	return res
}

func contains(strs []string, str string) bool {
	isContained := false
	for i := range strs {
		if strs[i] == str {
			isContained = true
			break
		}
	}

	return isContained
}
