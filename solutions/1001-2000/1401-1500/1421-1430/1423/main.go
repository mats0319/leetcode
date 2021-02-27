package mario

// 1 <= k <= cp.length
func maxScore(cardPoints []int, k int) int {
	minRemain := 0
	{
		cardRemain := 0
		for i := 0; i < len(cardPoints)-k; i++ {
			cardRemain += cardPoints[i]
		}

		minRemain = cardRemain
		for i := 0; i < k; i++ {
			cardRemain += cardPoints[i+len(cardPoints)-k] - cardPoints[i]
			if minRemain > cardRemain {
				minRemain = cardRemain
			}
		}
	}

	sum := 0
	for i := 0; i < len(cardPoints); i++ {
		sum += cardPoints[i]
	}

	return sum - minRemain
}
