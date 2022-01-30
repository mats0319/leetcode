package mario

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */
func guess(num int) int {
    if num == target {
        return 0
    } else if num < target {
        return 1
    } else {
        return -1
    }
}

var target int

func guessNumber(n int) int {
    res := 0

    left, right := 1, n
ALL:
    for left <= right {
        mid := left + (right-left)/2
        switch guess(mid) {
        case 0:
            res = mid
            break ALL
        case 1:
            left = mid+1
        case -1:
            right = mid-1
        }
    }

    return res
}
