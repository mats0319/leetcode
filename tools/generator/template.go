package main

import (
	"bytes"
	"go/format"
	"log"
	"text/template"
)

type testFileTemplateStruct struct {
	// params type in test case
	InputType  string
	ExpectType string

	InputParamsStruct    string // if input params' type more than 1, define struct
	InputParamsForInvoke string

	ExportedFuncName string // part of test function name, like 'Testxxx', without 'Test-' prefix
	InvokeFuncName   string // for invoke function
	SpelledFuncName  string // in error message
}

const (
	fileName     = "main.go"
	testFileName = "main_test.go"
)

const testFileTemplate = `
package mario

import "testing"

{{.InputParamsStruct}}

var testCase = []struct {
	In     {{.InputType}}
	Expect {{.ExpectType}}
}{
    // test cases here
	{},
}

func Test{{.ExportedFuncName}}(t *testing.T) {
	for i := range testCase {
		if {{.InvokeFuncName}}({{.InputParamsForInvoke}}) != testCase[i].Expect {
			t.Logf("{{.SpelledFuncName}} test failed on case: %d\n", i)
			t.Fail()
		}
	}
}
`

func (t *testFileTemplateStruct) fillTemplate() ([]byte, error) {
	templateIns, err := template.New("test file generate tool").Parse(testFileTemplate)
	if err != nil {
		log.Println("parse template failed, error: ", err)
		return nil, err
	}

	buffer := bytes.NewBufferString("")
	if err = templateIns.Execute(buffer, t); err != nil {
		log.Println("fill template failed, error: ", err)
		return nil, err
	}

	// format go code, unnecessary in generate no-Go code
	result, err := format.Source(buffer.Bytes())
	if err != nil {
		log.Println("format source failed, error: ", err)
		return nil, err
	}

	return result, nil
}
