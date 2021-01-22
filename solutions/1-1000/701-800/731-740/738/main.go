package mario

func monotoneIncreasingDigits(N int) int {
    if N% 10 == N {
        return N
    }

    digits := []int{N % 10}
    NCopy := N / 10
    isIncrease := true
    for NCopy != 0 {
        v := NCopy % 10
        digits = append([]int{v}, digits...)
        if v > digits[1] {
            isIncrease = false
        }
        NCopy /= 10
    }

    if isIncrease {
        return N
    }

    nineIndex := len(digits)
    for i := len(digits)-1; i > 0; i-- {
        if digits[i-1] > digits[i] {
            digits[i-1]--
            for j := i; j < nineIndex; j++ {
                digits[j] = 9
            }
            nineIndex = i
        }
    }

    res := 0
    for i := 0; i < len(digits); i++ {
        res = res * 10 + digits[i]
    }

    return res
}
