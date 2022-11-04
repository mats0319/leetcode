package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"os"
	"regexp"
)

func parseFuncDeclaration(fileName string) (funcName string, inputParams string, outputParams string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "open file failed"))
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatalln(errors.Wrap(err, "close file failed"))
		}
	}()

	fileScanner := bufio.NewScanner(file)
	for {
		if ok := fileScanner.Scan(); !ok {
			log.Fatalln("file scan finished, don't have a function to test") // file exit
		}

		var reEngine *regexp.Regexp
		reEngine, err = regexp.Compile(`^func (\w+)\((.*)\)(.*){$`) // e.g. 'func xxx(x int) error {', get 'xxx'/'x int'/'error'
		if err != nil {
			log.Fatalln(errors.Wrap(err, "init RE engine failed"))
		}

		matchedRes := reEngine.FindStringSubmatch(fileScanner.Text())
		if len(matchedRes) < 1 || len(matchedRes[0]) < 4 { // in res, first item is whole line, following are matched with '()'
			continue
		}

		funcName = matchedRes[1]
		inputParams = matchedRes[2]
		outputParams = matchedRes[3]

		break
	}

	return
}

func generateTemplateData(funcName string, inputParams string, outputParams string) *testFileTemplateStruct {
	result := &testFileTemplateStruct{
		ExpectType: outputParams, // require: only one output param without name, valid for leetcode

		ExportedFuncName: makeSureFirstCharBigCase(funcName),
		InvokeFuncName:   funcName,
		SpelledFuncName:  camelCaseToSpelledStyle(funcName),
	}

	paramList := parseInputParams(inputParams)
	if len(paramList) > 1 { // define input params struct
		result.InputType = "In"
		result.InputParamsStruct = paramsToStruct(paramList)
		result.InputParamsForInvoke = paramsToInput(paramList)
	} else if len(paramList) == 1 { // only one input param
		result.InputType = paramList[0].typ()
		result.InputParamsForInvoke = "testCase[i].In"
	}

	return result
}

// camelCaseToSpelledStyle make 'func name' from camel-case to spelled style
// e.g. 'FuncNameDemo' -> 'func name demo'
func camelCaseToSpelledStyle(word string) string {
	byteSlice := make([]byte, 0, len(word)*2)

	index := 0
	for i := 0; i < len(word); i++ {
		if 'A' <= word[i] && word[i] <= 'Z' {
			byteSlice = append(byteSlice, ' ')
			index++
			byteSlice = append(byteSlice, word[i]-'A'+'a')
		} else {
			byteSlice = append(byteSlice, word[i])
		}
	}

	byteSlice = bytes.TrimSpace(byteSlice) // trim all leading and trailing white space

	return string(byteSlice)
}

func paramsToStruct(params []*param) string {
	var structFields []byte
	for i := range params {
		field := makeSureFirstCharBigCase(params[i].string())
		structFields = append(structFields, fmt.Sprintf("%s\n", field)...)
	}

	return fmt.Sprintf("type In struct {\n%s}\n", string(structFields))
}

func paramsToInput(params []*param) string {
	var byteSlice []byte
	for i := range params {
		byteSlice = append(byteSlice, fmt.Sprintf(", testCase[i].In.%s", makeSureFirstCharBigCase(params[i].name))...)
	}

	if len(byteSlice) > 0 {
		byteSlice = byteSlice[2:]
	}

	return string(byteSlice)
}

func makeSureFirstCharBigCase(str string) string {
	if 'a' <= str[0] && str[0] <= 'z' {
		byteSlice := []byte(str)
		byteSlice[0] = byteSlice[0] - 'a' + 'A'
		str = string(byteSlice)
	}

	return str
}
