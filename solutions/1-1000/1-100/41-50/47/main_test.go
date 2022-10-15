package mario

import (
	"strconv"
	"testing"
)

var testCase = []struct {
	In     []int
	Expect [][]int
}{
	// test cases here
	{In: []int{2, 2, 1, 1}, Expect: [][]int{
		{1, 1, 2, 2},
		{1, 2, 1, 2},
		{1, 2, 2, 1},
		{2, 1, 1, 2},
		{2, 1, 2, 1},
		{2, 2, 1, 1},
	}},
}

func TestPermuteUnique(t *testing.T) {
	for i := range testCase {
		if !compareOnTwoDimensionIntSlice(permuteUnique(testCase[i].In), testCase[i].Expect) {
			t.Logf("permute unique test failed on case: %d\n", i)
			t.Fail()
		}
	}
}

func compareOnTwoDimensionIntSlice(a [][]int, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	aMap := make(map[string]int, len(a))
	for i := range a {
		aMap[serializeIntSlice(a[i])]++
		aMap[serializeIntSlice(b[i])]--
	}

	isEqual := true
	for _, value := range aMap {
		if value != 0 {
			isEqual = false
			break
		}
	}

	return isEqual
}

func serializeIntSlice(a []int) string {
	if len(a) < 1 {
		return ""
	}

	var (
		sep byte = ','
		res []byte
	)

	for i := range a {
		res = append(res, sep)
		res = append(res, strconv.Itoa(a[i])...)
	}

	return string(res[1:])
}
