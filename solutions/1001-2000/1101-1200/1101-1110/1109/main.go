package mario

func corpFlightBookings(bookings [][]int, n int) []int {
    res := make([]int, n+1)
    for _, v := range bookings {
        res[v[0]-1] += v[2]
        res[v[1]] -= v[2]
    }

    for i := 1; i < len(res)-1; i++ {
        res[i] += res[i-1]
    }

    return res[:len(res)-1]
}
