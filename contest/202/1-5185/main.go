package mario

func threeConsecutiveOdds(arr []int) bool {
    if len(arr) < 3 {
        return false
    }

    var oddsCount int

    for i := 0; i < len(arr) && oddsCount < 3; i++ {
        if arr[i] % 2 == 1 {
            oddsCount++
        } else {
            oddsCount = 0
        }
    }

    return oddsCount >= 3
}
