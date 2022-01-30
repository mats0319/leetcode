package main

const (
	fileName     = "main.go"
	testFileName = "main_test.go"
)

const (
	templateStr = `
package mario

import "testing"

{{.InputStructDefine}}

var testCase = []struct {
	In     {{.InputTypes}}
	Expect {{.ExpectType}}
}{
    // test cases here
	{},
}

func Test{{.FuncNameExported}}(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if {{.FuncNameActually}}({{.InputParams}}) != tcs[i].Expect {
			t.Errorf("{{.FuncNameSpelled}} test failed on case: %d\n", i)
		}
	}
}
`
)
