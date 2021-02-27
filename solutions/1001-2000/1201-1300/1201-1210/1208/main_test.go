package mario

import "testing"

type In struct {
	String0 string
	String1 string
	Int0    int
}

var testCase = []struct {
	In     In
	Expect int
}{
	// test cases here
	{In: In{"bcdf", "abcd", 3}, Expect: 3},
	{In: In{"cdef", "abcd", 3}, Expect: 1},
}

func TestEqualSubstring(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if equalSubstring(tcs[i].In.String1, tcs[i].In.String0, tcs[i].In.Int0) != tcs[i].Expect {
			t.Errorf("equal substring test failed on case: %d\n", i)
		}
	}
}
