package main

import (
	"bytes"
	"fmt"
	"strings"
)

func generateTemplateData(funcDeclaration string) *templateData {
	funcName, inputParams, outputParams := splitFuncDeclaration(funcDeclaration)
	result := &templateData{
		// require: only one output param without name, valid for leetcode
		ExpectType: outputParams,

		FuncActuallyName: funcName,
		FuncExportedName: makeFirstCharBigCase(funcName),
		FuncSpelledName:  generateFuncSpelledName(funcName),
	}

	result.formatInputs(inputParams)

	return result
}

func splitFuncDeclaration(funcDeclaration string) (funcName string, inputParams string, outputParams string) {
	firstSepIndex := strings.Index(funcDeclaration, "(")
	secondSepIndex := strings.Index(funcDeclaration, ")")

	funcName = trimAllSpace(funcDeclaration[:firstSepIndex])
	inputParams = removeVariableName(funcDeclaration[firstSepIndex+1 : secondSepIndex]) // already no space
	outputParams = trimAllSpace(funcDeclaration[secondSepIndex+1:])

	return
}

// removeVariableName make string like 'a int, b int' to 'int,int'
func removeVariableName(params string) string {
	paramsSlice := strings.Split(params, ",")
	for i := range paramsSlice {
		paramsSlice[i] = strings.TrimSpace(paramsSlice[i]) // trim leading and trailing white space
		paramsSlice[i] = paramsSlice[i][strings.Index(paramsSlice[i], " ")+1:]
	}

	return strings.Join(paramsSlice, ",")
}

// generateFuncSpelledName make 'func name' from camel-case to word-spell-case
// e.g. 'FuncNameCamelCase' -> 'func name camel case'
func generateFuncSpelledName(funcName string) string {
	byteSlice := make([]byte, 0, len(funcName)*2)

	index := 0
	for i := 0; i < len(funcName); i++ {
		if 'A' <= funcName[i] && funcName[i] <= 'Z' {
			byteSlice = append(byteSlice, ' ')
			index++
			byteSlice = append(byteSlice, funcName[i] - 'A' + 'a')
		} else {
			byteSlice = append(byteSlice, funcName[i])
		}
	}

	byteSlice = bytes.TrimSpace(byteSlice) // trim all leading and trailing white space

	return string(byteSlice)
}

// formatInputs 'inputs' is a brief input params str in func declaration, without params name and space
// e.g. 'int,int,bool,[]
func (fd *templateData) formatInputs(inputs string) {
	inputsSplit := strings.Split(inputs, ",")
	typeMap := make(map[string]int) // param type - appear times
	for i := range inputsSplit {
		typeMap[inputsSplit[i]]++
	}

	if len(typeMap) == 0 { // no input param
		return
	}

	if len(typeMap) == 1 { // only one type of input param, left 'InputStructDefine' empty
		for key := range typeMap {
			if typeMap[key] == 1 { // only one input param
				fd.InputTypes = key
				fd.InputParams = "testCase[i].In"
			} else { // several input params, in same type
				fd.InputTypes = "[]" + key

				inputParams := []byte("testCase[i].In[0]")
				for i := 1; i < typeMap[key]; i++ {
					inputParams = append(inputParams, fmt.Sprintf(", testCase[i].In[%d]", i)...)
				}

				fd.InputParams = string(inputParams)
			}
		}
	} else { // several types of input params, define 'input params struct'
		// name in test case
		fd.InputTypes = "In"

		// input params struct define
		var structFields []byte
		for key := range typeMap {
			for i := 0; i < typeMap[key]; i++ {
				structFields = append(structFields, typeToField(key, i)...)
			}
		}

		fd.InputStructDefine = fmt.Sprintf("type In struct {\n%s}\n", string(structFields))

		// input params when invoke func
		var inputParams []byte
		timesMap := make(map[string]int, len(typeMap)) // params - used times

		for i := range inputsSplit {
			index, _ := timesMap[inputsSplit[i]]

			inputParams = append(inputParams, ',', ' ')
			inputParams = append(inputParams, typeToInvoke(inputsSplit[i], index)...)

			timesMap[inputsSplit[i]]++
		}

		fd.InputParams = string(inputParams[2:])
	}

	return
}

// typeToField return a named struct field of 'type str'
// e.g.
// '[3]int' -> 'Int1Slice0 [3]int\n'
// '[][]int' -> 'Int2Slice0 [][]int\n'
// 'int' -> 'Int0 int\n'
func typeToField(typeStr string, index int) string {
	filedName := typeToFieldName(typeStr, index)
	return fmt.Sprintf("%s %s\n", filedName, typeStr)
}

// typeToInvoke return invoke name of 'type str'
// e.g.
// 'int' -> 'testCase[i].In.Int0'
// '[]int' -> 'testCase[i].In.Int1Slice0'
func typeToInvoke(typeStr string, index int) string {
	filedName := typeToFieldName(typeStr, index)
	return fmt.Sprintf("testCase[i].In.%s", filedName)
}

func typeToFieldName(typeStr string, index int) (fieldName string) {
	i := 0
	count := 0
	for i < len(typeStr) {
		if typeStr[i] != '[' {
			break
		}

		count++
		for i < len(typeStr) && typeStr[i] != ']' { // find next ']', for '[3]' / '[]' cases
			i++
		}

		i++
	}

	if count > 0 {
		fieldName = fmt.Sprintf("%s%dSlice%d", makeFirstCharBigCase(typeStr[i:]), count, index)
	} else {
		fieldName = fmt.Sprintf("%s%d", makeFirstCharBigCase(typeStr), index)
	}

	return
}
