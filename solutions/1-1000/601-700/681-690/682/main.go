package mario

import "strconv"

// ops is always valid
func calPoints(ops []string) int {
	scores := make([]int, len(ops))
	length := 0
	for _, s := range ops {
		switch s {
		case "+":
			scores[length] = scores[length-1] + scores[length-2]
			length++
		case "D":
			scores[length] = scores[length-1] * 2
			length++
		case "C":
			scores[length-1] = 0
			length--
		default: // integer
			value, _ := strconv.Atoi(s)
			scores[length] = value
			length++
		}
	}

	return sum(scores[:length])
}

func sum(src []int) (res int) {
	for _, v := range src {
		res += v
	}

	return
}
