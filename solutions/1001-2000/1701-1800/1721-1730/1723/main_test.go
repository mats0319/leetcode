package mario

import "testing"

type In struct {
	IntSlice0 []int
	Int0      int
}

var testCase = []struct {
	In     In
	Expect int
}{
	// test cases here
	{In{[]int{1, 2, 4, 7, 8}, 2}, 11},
	{In{[]int{11, 2, 4}, 1}, 17},
	{In{[]int{5, 5, 4, 4, 4}, 2}, 12},
	//{In{[]int{6518448, 8819833, 7991995, 7454298, 2087579, 380625, 4031400,
	//	2905811, 4901241, 8480231, 7750692, 3544254}, 4}, 16274131},
}

func TestMinimumTimeRequired(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if minimumTimeRequired(tcs[i].In.IntSlice0, tcs[i].In.Int0) != tcs[i].Expect {
			t.Errorf("minimum time required test failed on case: %d\n", i)
		}
	}
}
