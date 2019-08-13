package ri

import "testing"

var testCase = []int{
	123,
	-123,
	120,
	1534236469,
}

var result = []int{
	321,
	-321,
	21,
	0,
}

func TestReverse(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if reverse(tcs[i]) != result[i] {
			t.Errorf("reverse integer test failed on case: %d", i)
		}
	}
}
