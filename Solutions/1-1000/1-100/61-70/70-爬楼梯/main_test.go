package mario

import "testing"

var testCase = []struct{
    In int
    Expect int
}{
    {2, 2},
    {3, 3},
    {4, 5},
}

func TestClimbStairs(t *testing.T) {
    tcs := testCase
    for i := range tcs {
        if climbStairs(tcs[i].In) != tcs[i].Expect {
            t.Errorf("climb stairs test failed on case: %d", i)
        }
    }
}
