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

func Test{{.FuncExportedName}}(t *testing.T) {
	for i := range testCase {
		if {{.FuncActuallyName}}({{.InputParams}}) != testCase[i].Expect {
			t.Logf("{{.FuncSpelledName}} test failed on case: %d\n", i)
			t.Fail()
		}
	}
}
`
)
