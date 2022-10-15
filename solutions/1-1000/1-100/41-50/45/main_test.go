package mario

import "testing"

var testCase = []struct {
    In     []int
    Expect int
}{
    // test cases here
    {In: []int{2, 0, 2, 0, 1}, Expect: 2},
    {In: []int{2, 3, 1, 1, 4}, Expect: 2},
    {In: []int{2, 3, 0, 1, 4}, Expect: 2},
    {In: []int{0}, Expect: 0},
    {In: []int{1}, Expect: 0},
}

func TestJump(t *testing.T) {
    for i := range testCase {
        if jump(testCase[i].In) != testCase[i].Expect {
            t.Logf("jump test failed on case: %d\n", i)
            t.Fail()
        }
    }
}
