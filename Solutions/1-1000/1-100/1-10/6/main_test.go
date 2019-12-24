package mario

import (
	"testing"
)

type In struct {
	Str string
	Int int
}

var testCase = []struct {
	In     In
	Expect string
}{
	{In{"PAYPALISHIRING", 3}, "PAHNAPLSIIGYIR"},
	{In{"PAYPALISHIRING", 4}, "PINALSIGYAHRPI"},
}

func TestConvert(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if convert(tcs[i].In.Str, tcs[i].In.Int) != tcs[i].Expect {
			t.Errorf("zig zag conversion test failed on case: %d", i)
		}
	}
}

func TestConvert2(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if convert2(tcs[i].In.Str, tcs[i].In.Int) != tcs[i].Expect {
			t.Errorf("zig zag conversion solution: 2 test failed on case: %d", i)
		}
	}
}
