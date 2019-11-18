package main

import "testing"

var testCase = []struct {
	In     string
	Except int
}{
	{"42", 42},
	{"   -42", -42},
	{"4193 with words", 4193},
	{"words and 987", 0},
	{"-91283472332", -2147483648},
	{" ", 0},
	{"   ", 0},
	{"9223372036854775808", 2147483647},
}

func TestMyAtoi(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if myAtoi(tcs[i].In) != tcs[i].Except {
			t.Errorf("string to integer test failed on case: %d", i)
		}
	}
}
