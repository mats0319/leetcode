package sort

import (
	"testing"
)

var testCase = []struct {
	In     []int
	Expect []int
}{
	{[]int{2, 6, 7, 4, 1, 3, 8, 9, 5}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
	{[]int{3, 4, 2, 1, 5, 6, 7, 8}, []int{1, 2, 3, 4, 5, 6, 7, 8}},
}

func TestBubbleSort(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		BubbleSort(tcs[i].In)
		if !compareOnIntSlice(tcs[i].In, tcs[i].Expect) {
			t.Errorf("bubble sort test failed on case: %d", i)
		}
	}
}

func TestSelectionSort(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		SelectionSort(tcs[i].In)
		if !compareOnIntSlice(tcs[i].In, tcs[i].Expect) {
			t.Errorf("selection sort test failed on case: %d", i)
		}
	}
}

func TestMergeSort(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		MergeSort(tcs[i].In)
		if !compareOnIntSlice(tcs[i].In, tcs[i].Expect) {
			t.Errorf("merge sort test failed on case: %d", i)
		}
	}
}

func TestQuickSort(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		QuickSort(tcs[i].In)
		if !compareOnIntSlice(tcs[i].In, tcs[i].Expect) {
			t.Errorf("quick sort test failed on case: %d", i)
		}
	}
}

func TestCountingSort(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		CountingSort(tcs[i].In)
		if !compareOnIntSlice(tcs[i].In, tcs[i].Expect) {
			t.Errorf("counting sort test failed on case: %d", i)
		}
	}
}

func TestInsertionSort(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		InsertionSort(tcs[i].In)
		if !compareOnIntSlice(tcs[i].In, tcs[i].Expect) {
			t.Errorf("insertion sort test failed on case: %d", i)
		}
	}
}

func compareOnIntSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
