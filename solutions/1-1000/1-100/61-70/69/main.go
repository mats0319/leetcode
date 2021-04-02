package mario

func mySqrt(x int) int {
    res := 0
    for (res+1) * (res+1) <= x {
        res++
    }

    return res
}
