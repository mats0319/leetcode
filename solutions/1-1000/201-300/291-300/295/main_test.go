package mario

import (
    "fmt"
    "testing"
)

func Test295(t *testing.T) {
    case1()
}

func case1() {
    res := make([]string, 0)

    ins := Constructor()
    ins.AddNum(1)
    ins.AddNum(2)
    res = addResult(res, ins.FindMedian())
    ins.AddNum(3)
    res = addResult(res, ins.FindMedian())

    fmt.Println(len(res), res)
}

func addResult(arr []string, item float64) []string {
    return append(arr, fmt.Sprintf("%.4f", item))
}
