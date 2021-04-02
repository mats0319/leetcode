package mario

import (
	"strconv"
	"strings"
)

func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}

	res := ""
	{
		flag := 1
		if numerator < 0 {
			flag *= -1
			numerator *= -1
		}
		if denominator < 0 {
			flag *= -1
			denominator *= -1
		}

		if flag < 0 {
			res = "-"
		}
	}

	div := numerator / denominator
	numerator -= denominator * div

	res += strconv.Itoa(div)

	if numerator == 0 {
		return res
	}

	res += "."
	{
		numerator *= 10
		exps := make([]int, 0) // keep next bei4_chu2_shu4
		for numerator != 0 {
			exps = append(exps, numerator)

			div = numerator / denominator
			res += strconv.Itoa(div)

			numerator -= denominator * div
			numerator *= 10

			if index := isLoop(exps, numerator); index >= 0 {
				loopStart := strings.Index(res, ".") + 1 + index
				res = res[:loopStart] + "(" + res[loopStart:] + ")"

				break
			}

		}
	}

	return res
}

func isLoop(array []int, value int) int {
	index := -1
	for i := 0; i < len(array); i++ {
		if value == array[i] {
			index = i
			break
		}
	}

	return index
}
