package motsa

import (
	"testing"
)

type IntSliceGroup struct {
	IntSliceOne []int
	IntSliceTwo []int
}

var testCase = []IntSliceGroup{
	{[]int{1, 2}, []int{3, 4}},
	{[]int{1, 3}, []int{2}},
	{[]int{1, 4}, []int{2, 3}},
	{[]int{1}, []int{2, 3}},
}

var result = []float64{
	2.5,
	2.0,
	2.5,
	2,
}

func TestFindMedianSortedArrays(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if findMedianSortedArrays(tcs[i].IntSliceOne, tcs[i].IntSliceTwo) != result[i] {
			t.Errorf("find median sorted arrays test failed on case: %d", i)
		}
	}
}

func TestFindMedianSortedArrays2(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if findMedianSortedArrays2(tcs[i].IntSliceOne, tcs[i].IntSliceTwo) != result[i] {
			t.Errorf("find median sorted arrays test failed on case: %d", i)
		}
	}
}
