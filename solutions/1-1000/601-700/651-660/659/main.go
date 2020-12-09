package mario

func isPossible(nums []int) bool {
    left := make(map[int]int, len(nums))
    for _, v := range nums {
        left[v]++
    }

    endCount := make(map[int]int, 0)
    possibleFlag := true
    for _, v := range nums {
        if left[v] == 0 {
            continue
        }

        if endCount[v-1] > 0 {
            left[v]--
            endCount[v-1]--
            endCount[v]++
        } else if left[v+1] > 0 && left[v+2] > 0 {
            left[v]--
            left[v+1]--
            left[v+2]--
            endCount[v+2]++
        } else {
            possibleFlag = false
            break
        }
    }

    return possibleFlag
}
