package mario

import (
    "sort"
    "strconv"
    "testing"
)

type In struct {
    Int1Slice0 []int
    Int0       int
}

var testCase = []struct {
    In     In
    Expect [][]int
}{
    // test cases here
    {In: In{Int1Slice0: []int{1, 2, 3, 4, 5}, Int0: 5},
       Expect: [][]int{{1, 1, 1, 1, 1}, {1, 1, 1, 2}, {1, 1, 3}, {1, 4}, {1, 2, 2}, {2, 3}, {5}}},
    {In: In{Int1Slice0: []int{1, 2, 3, 5, 6, 7}, Int0: 9},
        Expect: [][]int{
            {1, 1, 1, 1, 1, 1, 1, 1, 1},
            {1, 1, 1, 1, 1, 1, 1, 2},
            {1, 1, 1, 1, 1, 1, 3},
            {1, 1, 1, 1, 1, 2, 2},
            {1, 1, 1, 1, 2, 3},
            {1, 1, 1, 1, 5},
            {1, 1, 1, 2, 2, 2},
            {1, 1, 1, 3, 3},
            {1, 1, 1, 6},
            {1, 1, 2, 2, 3},
            {1, 1, 2, 5},
            {1, 1, 7},
            {1, 2, 2, 2, 2},
            {1, 2, 3, 3},
            {1, 2, 6},
            {1, 3, 5},
            {2, 2, 2, 3},
            {2, 2, 5},
            {2, 7},
            {3, 3, 3},
            {3, 6},
        },
    },
}

func TestCombinationSum(t *testing.T) {
    for i := range testCase {
        if !compareOnTwoDimensionIntSliceNotStrict(combinationSum(testCase[i].In.Int1Slice0, testCase[i].In.Int0), testCase[i].Expect) {
            t.Logf("combination sum test failed on case: %d\n", i)
            t.Fail()
        }
    }
}

// compareOnTwoDimensionIntSliceNotStrict 比较两个二维int数组，与一维、二维上的顺序都无关
func compareOnTwoDimensionIntSliceNotStrict(a [][]int, b [][]int) bool {
    if len(a) != len(b) {
        return false
    }

    aMap := make(map[string]int, len(a))
    for i := range a {
        aMap[intSliceToStringByOrder(a[i])]++
        aMap[intSliceToStringByOrder(b[i])]--
    }

    isEqual := true
    for _, value := range aMap {
        if value != 0 {
            isEqual = false
            break
        }
    }

    return isEqual
}

func intSliceToStringByOrder(slice []int) string {
    sort.Ints(slice)

    var res []byte
    for i := range slice {
        res = append(res, strconv.Itoa(slice[i])...)
    }

    return string(res)
}
