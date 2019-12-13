package mario

import (
    "testing"
)

var testCase = []struct {
    In []int
    Expect int
}{
    {[]int{7, 3}, 28},
    {[]int{3, 2}, 3},
}

func TestUniquePaths(t *testing.T) {
    tcs := testCase
    for i := range tcs {
        if uniquePaths(tcs[i].In[0], tcs[i].In[1]) != tcs[i].Expect {
            t.Errorf("unique paths test failed on case: %d", i)
        }
    }
}
