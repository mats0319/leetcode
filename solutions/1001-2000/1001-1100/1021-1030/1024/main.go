package mario

func videoStitching(clips [][]int, T int) int {
	count := 0
	for end, index := 0, len(clips)-1; end < T; count++ {
		index, end = nextInterval(clips, index, end)

		if end == -1 {
			count = -1
			break
		}
	}

	return count
}

func nextInterval(src [][]int, index, limit int) (int, int) {
	end := limit
	for i := 0; i <= index; i++ {
		if src[i][0] <= limit {
			if src[i][1] > end {
				end = src[i][1]
			}

			src[i], src[index] = src[index], src[i]
			index--
			i--
		}
	}

	// no new longer interval or src is nil
	if end == limit {
		end = -1
	}

	return index, end
}
