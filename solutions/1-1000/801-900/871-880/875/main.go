package mario

func minEatingSpeed(piles []int, h int) int {
	left := 1
	right := 0
	for i := range piles {
		if piles[i] > right {
			right = piles[i]
		}
	}

	minSpeed := 0
	for left <= right {
		minSpeed = left + (right-left)/2

		currTime := 0
		for i := 0; i < len(piles) && currTime <= h; i++ {
			currTime += (piles[i] + minSpeed - 1) / minSpeed
		}

		if currTime > h {
			left = minSpeed + 1
		} else {
			right = minSpeed - 1
		}
	}

	return minSpeed
}
