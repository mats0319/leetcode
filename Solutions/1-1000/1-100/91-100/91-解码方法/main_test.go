package mario

import "testing"

var testCase = []struct{
    In string
    Expect int
}{
    {"12", 2},
    {"226", 3},
    {"0", 0},
    {"00", 0},
}

func TestNumDecodings(t *testing.T) {
    tcs := testCase
    for i := range tcs {
        if numDecodings(tcs[i].In) != tcs[i].Expect {
            t.Errorf("num decodings test failed on case: %d", i)
        }
    }
}
