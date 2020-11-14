package mario

import "sort"

var bit = make([]int, 10001)

func init() {
    for i := 0; i < 10001; i++ {
        bit[i] = bit[i/2] + i % 2
    }
}

func sortByBits(arr []int) []int {
    sort.Slice(arr, func(i, j int) bool {
        x, y := arr[i], arr[j]
        cx, cy := bit[x], bit[y]
        return cx < cy || cx == cy && x < y
    })

    return arr
}
