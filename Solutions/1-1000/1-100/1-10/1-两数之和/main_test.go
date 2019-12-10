package mario

import "testing"

var testCase = []struct {
	In     []int
	Expect int
}{
	{[]int{2, 7, 11, 13}, 9},
}

func TestTwoSum(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if !verify(tcs[i].In, twoSum(tcs[i].In, tcs[i].Expect), tcs[i].Expect) {
			t.Errorf("two sum test failed on case: %d", i)
		}
	}
}

func verify(src, index []int, target int) bool {
	var result int
	for i := range index {
		result += src[index[i]]
	}
	return result == target
}
