package mario

import "testing"

type In struct {
	IntSlice0 [][]int
	Int0      int
}

var testCase = []struct {
	In     In
	Expect []int
}{
	// test cases here
	{In{[][]int{{1, 2, 10}, {2, 3, 20}, {2, 5, 25}}, 5}, []int{10, 55, 45, 25, 25}},
}

func TestCorpFlightBookings(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if !compareOnIntSlice(corpFlightBookings(tcs[i].In.IntSlice0, tcs[i].In.Int0), tcs[i].Expect) {
			t.Errorf("corp flight bookings test failed on case: %d\n", i)
		}
	}
}

func compareOnIntSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	isEqual := true
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			isEqual = false
			break
		}
	}

	return isEqual
}
