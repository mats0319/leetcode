package main

import (
	"testing"
)

type In struct {
	Is     []int
	Target int
}

var testCase = []struct {
	In     In
	Except [][]int
}{
	{In{[]int{-3, -2, -1, 0, 0, 1, 2, 3}, 0},
		[][]int{{-3, -2, 2, 3}, {-3, -1, 1, 3}, {-3, 0, 0, 3}, {-3, 0, 1, 2}, {-2, -1, 0, 3}, {-2, -1, 1, 2},
			{-2, 0, 0, 2}, {-1, 0, 0, 1}}},
}

func TestFourSum(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if !compareOnErWeiInt(fourSum(tcs[i].In.Is, tcs[i].In.Target), tcs[i].Except) {
			t.Errorf("four sum test failed on case: %d", i)
		}
	}
}

func compareOnErWeiInt(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	var (
		bMap  = make(map[int][]int, len(b))
		index int
	)

	for i := range b {
		bMap[i] = b[i]
	}

	for i := range a {
		index = contain(bMap, a[i])
		if index != -1 {
			delete(bMap, index)
		} else {
			return false
		}
	}

	return true
}

// if map contain 'is', return map.key
// or return -1 means not contain
func contain(m map[int][]int, is []int) int {
	var (
		isFlag  []int
		isMatch bool
	)

	for i := range m {
		isFlag = make([]int, len(is))

		for j := range m[i] {
			for k := range is {
				if m[i][j] == is[k] && isFlag[k] == 0 { // match and not repeated
					isFlag[k] = 1
					break
				}
			}
		}

		isMatch = true
		for j := range isFlag {
			if isFlag[j] != 1 {
				isMatch = false
				break
			}
		}

		if isMatch {
			return i
		}
	}

	return -1
}
