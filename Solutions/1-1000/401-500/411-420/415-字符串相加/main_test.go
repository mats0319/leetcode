package main

import (
	"testing"
)

type In struct {
	Num1 string
	Num2 string
}

var testCase = []struct {
	In     In
	Except string
}{
	{In{"123", "4567"}, "4690"},
	{In{"0", "0"}, "0"},
	{In{"9", "99"}, "108"},
	{In{"584", "18"}, "602"},
}

func TestAddStrings(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if addStrings(tcs[i].In.Num1, tcs[i].In.Num2) != tcs[i].Except {
			t.Errorf("add strings test failed on case: %d", i)
		}
	}
}
