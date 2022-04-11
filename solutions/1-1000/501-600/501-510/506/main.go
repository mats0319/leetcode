package mario

import "strconv"

func findRelativeRanks(score []int) []string {
    length := len(score)

    heap := NewBinaryMinimumHeap(score...)
    valueRankMap := make(map[int]string, length)
    for i := length; ; i-- {
        v, err := heap.Pop()
        if err != nil {
            break
        }

        switch i {
        case 1:
            valueRankMap[v] = "Gold Medal"
        case 2:
            valueRankMap[v] = "Silver Medal"
        case 3:
            valueRankMap[v] = "Bronze Medal"
        default:
            valueRankMap[v] = strconv.Itoa(i)
        }
    }

    res := make([]string, 0, length)
    for i := range score {
        res = append(res, valueRankMap[score[i]])
    }

    return res
}
