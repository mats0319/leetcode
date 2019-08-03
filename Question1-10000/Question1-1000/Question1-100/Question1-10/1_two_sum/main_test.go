package ts

import "testing"

var testCase = [][]int{
    {2, 7, 11, 13},
}

var result = []int{
    9,
}

func TestTwoSum(t *testing.T) {
    tcs := testCase
    for i := range tcs {
        if !verify(testCase[i], twoSum(tcs[i], result[i]), result[i]) {
            t.Errorf("two sum test failed on case: %d", i)
        }
    }
}

func verify(src, index []int, target int) bool {
    var result int
    for i := range index {
        result += src[index[i]]
    }
    return result == target
}
