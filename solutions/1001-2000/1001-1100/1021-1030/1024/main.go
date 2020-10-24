package mario

func videoStitching(clips [][]int, T int) int {
    end := 0
    count := 0
    for end < T && end != -1 {
        clips, end = nextInterval(clips, end)
        count++
    }

    if end == -1 {
        count = -1
    }

    return count
}

func nextInterval(src [][]int, limit int) ([][]int, int) {
    index := len(src)-1
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

    return src[:index+1], end
}
