package mario

import "testing"

type Params struct {
	I, J int
}

var testCase = []struct {
	In     []int
	Params []Params
	Expect []int
}{
	// test cases here
	{[]int{-2, 0, 3, -5, 2, -1}, []Params{{0, 2}, {2, 5}, {0, 5}}, []int{1, -1, -3}},
}

func TestSumRange(t *testing.T) {
	tcs := testCase
	var structure *NumArray = new(NumArray)
	for i := range tcs {
		if len(tcs[i].Params) != len(tcs[i].Expect) {
			t.Errorf("test case %d's params and expect are not matched on quantity\n", i)
		}
		for j := range tcs[i].Params {
			*structure = Constructor(tcs[i].In)
			if structure.SumRange(tcs[i].Params[j].I, tcs[i].Params[j].J) != tcs[i].Expect[j] {
				t.Errorf("sum range test failed on case: %d, %d\n", i, j)
			}
		}
	}
}
