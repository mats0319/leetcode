package mario

func candy(ratings []int) int {
	if len(ratings) < 2 {
		return len(ratings)
	}

	candyNum := make([]int, len(ratings))
	for i := range candyNum {
		candyNum[i] = 1
	}

	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			candyNum[i] = candyNum[i-1] + 1
		}
	}

	sum := candyNum[len(ratings)-1]
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			candyNum[i] = max(candyNum[i], candyNum[i+1]+1)
		}

		sum += candyNum[i]
	}

	return sum
}

func max(a, b int) (res int) {
	if a > b {
		res = a
	} else {
		res = b
	}

	return
}
