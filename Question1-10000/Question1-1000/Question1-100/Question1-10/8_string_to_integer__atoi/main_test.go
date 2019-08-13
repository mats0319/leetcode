package sti

import "testing"

var testCase = []string{
	"42",
	"   -42",
	"4193 with words",
	"words and 987",
	"-91283472332",
	" ",
	"   ",
	"9223372036854775808",
}

var result = []int{
	42,
	-42,
	4193,
	0,
	-2147483648,
	0,
	0,
	2147483647,
}

func TestMyAtoi(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if myAtoi(tcs[i]) != result[i] {
			t.Errorf("string to integer test failed on case: %d", i)
		}
	}
}
