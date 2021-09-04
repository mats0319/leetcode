package mario

var square []int = make([]int, 0, 101) // [0, 1, 4, 9, ..., 10000]

// 1 <= n <= 10^4
func numSquares(n int) int {
	initM() // for leetcode

	index, equal := getIndex(n)
	if equal {
		return 1
	}

	return minSummary(n, index)
}

func initM() {
	if len(square) >= 100 {
		return
	}

	for i := 0; i < cap(square); i++ {
		square = append(square, i*i)
	}

	return
}

func getIndex(n int) (index int, equal bool) {
	left := 1
	right := len(square) - 1
	for left <= right {
		mid := left + (right-left)/2
		if n == square[mid] || n == square[mid-1] {
			equal = true
			break
		} else if square[mid-1] < n && n < square[mid] {
			index = mid
			break
		} else if n > square[mid] {
			left = mid + 1
		} else if n < square[mid-1] {
			right = mid - 1
		}
	}

	return
}

func minSummary(target, index int) int {
	dp := make([]int, target+1)
	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		dp[i] = i
		for j := 1; j < index && i >= square[j]; j++ {
			if dp[i] > dp[i-square[j]]+1 {
				dp[i] = dp[i-square[j]] + 1
			}
		}
	}

	return dp[target]
}
