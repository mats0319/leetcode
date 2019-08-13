package zzc

import (
	"testing"
)

type TestCase struct {
	Str string
	Int int
}

var testCase = []TestCase{
	{"PAYPALISHIRING", 3},
	{"PAYPALISHIRING", 4},
}

var result = []string{
	"PAHNAPLSIIGYIR",
	"PINALSIGYAHRPI",
}

func TestConvert(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if convert(tcs[i].Str, tcs[i].Int) != result[i] {
			t.Errorf("zig zag conversion test failed on case: %d", i)
		}
	}
}

func TestConvert2(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if convert2(tcs[i].Str, tcs[i].Int) != result[i] {
			t.Errorf("zig zag conversion solution: 2 test failed on case: %d", i)
		}
	}
}
