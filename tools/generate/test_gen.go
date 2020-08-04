// Test_gen can generate universal leetcode test file,
// input a question's ID with flag '-n', and we will generate the test file and redirect to it's position.
//
// Once more, We support universal questions only, some special questions are not support, for example:
// - return value's type can't use '==' to judge equal
// - param type is double-level slice
//
// Also, it is still just a template, you need to add test cases yourself.
// -_- ... by listing the defects, I doubt if I write this is meaningful?
// just for fun, I write it for studying mainly. ^_^
//
// Mario, 2019.12.24
package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"github.com/pkg/errors"
	"go/format"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type FunctionDeclaration struct {
	InputTypes string
	ExpectType string

	InputStructDefine string
	InputParams       string

	FuncNameExported string
	FuncNameActually string
	FuncNameSpelled  string
}

const (
	fileName     = "main.go"
	testFileName = "main_test.go"

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

var (
	help       bool
	questionId int
)

func init() {
	flag.BoolVar(&help, "h", false, "this help")
	flag.IntVar(&questionId, "n", -1, "give out question number\n"+
		"we will read function declaration and redirect generated file automatically\n"+
		"a valid number is from 1 to 9999.")

	flag.Parse()

	if help {
		log.Println("Options: ")
		flag.PrintDefaults()
		os.Exit(0)
	}

	if questionId == -1 {
		log.Fatalln("Please input a valid question's ID, more information in help with '-h' flag. ")
	}
}

func main() {
	log.Println("test file generate tool started. ")
	defer log.Println("test file generate tool finished. ")

	generateFile(replaceTemplate(formatFuncDeclaration(generateDir(questionId))))
}

func generateDir(num int) string {
	var result string
	var tag int
	if num%10 == 0 {
		tag = -1
	}
	result = fmt.Sprintf("%d-%d/", (num/10+tag)*10+1, (num/10+tag)*10+10)
	for i := 100; i < 10000; i *= 10 {
		result = fmt.Sprintf("%d-%d/", num/i*i+1, num/i*i+i) + result
	}

	return fmt.Sprintf("../../Solutions/%s%d/", result, num)
}

func formatFuncDeclaration(dir string) (*FunctionDeclaration, string) {
	formatDeclaration := splitAndTrimAllSpace(getFuncDeclaration(dir)) // fd: func name + inputs + output
	result := &FunctionDeclaration{
		ExpectType:       formatDeclaration[2],
		FuncNameActually: formatDeclaration[0],
	}

	result.formatInputs(formatDeclaration[1])
	result.formatFuncNames(formatDeclaration[0])

	return result, dir
}

func getFuncDeclaration(dir string) string {
	file, err := os.OpenFile(dir+fileName, os.O_RDONLY, 0744)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "open file failed"))
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatalln(errors.Wrap(err, "close file failed"))
		}
	}()

	fileScanner := bufio.NewScanner(file)
	for {
		if ok := fileScanner.Scan(); !ok {
			log.Fatalln("file scan finished, don't have a function to test")
		}
		if len(fileScanner.Text()) > 5 && fileScanner.Text()[:5] == "func " { // todo: RE replace
			break
		}
	}

	return fileScanner.Text()[5 : len(fileScanner.Text())-1]
}

func splitAndTrimAllSpace(funcDeclaration string) (result []string) {
	var t int
	for t = 0; t < len(funcDeclaration) && funcDeclaration[t] != '('; t++ {
	}
	result = append(result, funcDeclaration[:t]) // func name

	for i := t + 1; i < len(funcDeclaration); i++ {
		if funcDeclaration[i] == ')' {
			result = append(result, funcDeclaration[t+1:i]) // input param(s)
			result = append(result, funcDeclaration[i+1:]) // output param(s)
			break
		}
	}

	result[1] = removeVariableName(result[1])
	// todo: remove return value's name, if exist, necessary?

	for i := 0; i <= 2; i += 2 {
		for j := 0; j < len(result[i]); j++ {
			if result[i][j] == ' ' {
				result[i] = result[i][:j] + result[i][j+1:]
				j--
			}
		}
	}

	return
}

func removeVariableName(src string) string {
	srcSplit := strings.Split(src, ",")
	for i := range srcSplit {
		srcSplit[i] = strings.TrimSpace(srcSplit[i])
		srcSplit[i] = srcSplit[i][strings.Index(srcSplit[i], " ")+1:]
	}

	return strings.Join(srcSplit, ",")
}

func (fd *FunctionDeclaration) formatInputs(inputs string) {
	inputsSplit := strings.Split(inputs, ",")
	m := make(map[string]int)
	for i := range inputsSplit {
		m[inputsSplit[i]]++
	}

	if len(m) == 0 { // default is ok
	} else if len(m) == 1 {
		for key := range m {
			fd.InputTypes = key

			if m[key] == 1 {
				fd.InputParams = "tcs[i].In"
			} else { // m[key] > 1
				fd.InputTypes = "[]" + fd.InputTypes

				for i := 0; i < m[key]; i++ {
					fd.InputParams = fmt.Sprintf("%s, tcs[i].In[%s]", fd.InputParams, strconv.Itoa(i))
				}
				fd.InputParams = fd.InputParams[2:]
			}
		}
	} else { // len(m) > 1
		fd.InputTypes = "In"

		for key := range m {
			for i := 0; i < m[key]; i++ {
				if 'a' <= key[0] && key[0] <= 'z' {
					fd.InputStructDefine = fmt.Sprintf("%s%s%s%d %s\n", fd.InputStructDefine, string(key[0]-'a'+'A'), key[1:], i, key)
				} else { // key[0] = '['
					fd.InputStructDefine = fmt.Sprintf("%s%s%sSlice%d %s\n", fd.InputStructDefine, string(key[2]-'a'+'A'), key[3:], i, key)
				}
			}
		}
		fd.InputStructDefine = fmt.Sprintf("type In struct {\n%s}\n", fd.InputStructDefine)

		for i := range inputsSplit {
			if 'a' <= inputsSplit[i][0] && inputsSplit[i][0] <= 'z' {
				fd.InputParams = fmt.Sprintf("%s, tcs[i].In.%s%s%d", fd.InputParams, string(inputsSplit[i][0]-'a'+'A'), inputsSplit[i][1:], m[inputsSplit[i]]-1)
			} else { // param is a slice
				fd.InputParams = fmt.Sprintf("%s, tcs[i].In.%s%sSlice%d", fd.InputParams, string(inputsSplit[i][2]-'a'+'A'), inputsSplit[i][3:], m[inputsSplit[i]]-1)
			}
			m[inputsSplit[i]]--
		}
		fd.InputParams = fd.InputParams[2:]
	}

	return
}

func (fd *FunctionDeclaration) formatFuncNames(funcName string) {
	if 'a' <= funcName[0] && funcName[0] <= 'z' {
		funcName = string(funcName[0]+'A'-'a') + funcName[1:]
	}
	fd.FuncNameExported = funcName

	var (
		funcNameSplit = make([]byte, len(funcName)*2)
		t             int
	)
	for i := 0; i < len(funcName); i, t = i+1, t+1 {
		if 'A' <= funcName[i] && funcName[i] <= 'Z' {
			funcNameSplit[t] = ' '
			t++
			funcNameSplit[t] = funcName[i] - 'A' + 'a'
		} else {
			funcNameSplit[t] = funcName[i]
		}
	}
	fd.FuncNameSpelled = string(funcNameSplit[1:t])

	return
}

func replaceTemplate(fd *FunctionDeclaration, dir string) ([]byte, string) {
	t, err := template.New("test file generate tool").Parse(templateStr)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "parse template failed"))
	}

	buff := bytes.NewBufferString("")
	if err = t.Execute(buff, fd); err != nil {
		log.Fatalln(errors.Wrap(err, "fill template failed"))
	}

	result, err := format.Source(buff.Bytes())
	if err != nil {
		log.Fatalln(errors.Wrap(err, "format source failed"))
	}

	return result, dir
}

func generateFile(src []byte, dir string) {
	fullName := dir + testFileName
	if _, err := os.Stat(fullName); os.IsNotExist(err) {
		if err = ioutil.WriteFile(fullName, src, 0755); err != nil {
			log.Fatalln(errors.Wrap(err, "write failed"))
		}
	} else {
		log.Fatalln("file exist with name: ", testFileName)
	}

	return
}
