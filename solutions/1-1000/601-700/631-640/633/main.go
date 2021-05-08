package mario

import "math"

func judgeSquareSum(c int) bool {
	left := 0
	right := int(math.Sqrt(float64(c)))
	isValid := false
	for left <= right {
		sum := left*left + right*right
        if sum == c {
            isValid = true
            break
        } else if sum > c {
            right--
        } else {
            left++
        }
	}

	return isValid
}
