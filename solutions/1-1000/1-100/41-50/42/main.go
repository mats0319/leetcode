package mario

func trap(height []int) int {
    if len(height) < 3 {
        return 0
    }

    res := 0
    for i := 1; i < len(height)-1; i++ {
        left := findMax(height, 0, i-1)
        right := findMax(height, i+1, len(height)-1)

        h := min(left, right)-height[i]
        if h > 0 {
            res += h
        }
    }

    return res
}

// make sure 'start' & 'end' are valid index
func findMax(slice []int, start, end int) (max int) {
    for i := start; i <= end; i++ {
        if slice[i] > max {
            max = slice[i]
        }
    }

    return max
}

func min(a, b int) (res int) {
    if a < b {
        res = a
    } else {
        res = b
    }

    return
}
