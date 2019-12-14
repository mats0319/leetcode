package mario

import "testing"

var testCase = []struct {
    In []int
    Expect int
}{
    {[]int{7,1,5,3,6,4}, 5},
    {[]int{7,6,4,3,1}, 0},
}

func TestMaxProfit(t *testing.T) {
    tcs := testCase
    for i := range tcs {
        if maxProfit(tcs[i].In) != tcs[i].Expect {
            t.Errorf("max profit test failed on case: %d", i)
        }
    }
}
