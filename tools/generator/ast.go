package main

import (
	"strings"
)

type token struct {
	value string
	typ   string
}

type param struct {
	name      string
	baseType  string
	isArray   bool   // describe on 'param'
	arrayFlag string // e.g. `[3]`/`[][]`
	isPointer bool   // describe on 'base type', e.g. `[]*int` also consider as pointer
}

const (
	TokenType_VariableName = "variable_name"
	TokenType_VariableType = "variable_type"
)

const (
	ScanStatus_ReadingName = iota
	ScanStatus_ReadingType
)

func parseInputParams(str string) []*param {
	return parseTokens(scan(str))
}

// scan rule: e.g. ` a [3]string , b []*int ` -> {a, name}, {[3]string, type}, {b, name}, {[]*int, type}
// blank:
//   0: index < 0: skip
//      index >= 0: [index, i) is 'name', status -> 1, index = -1
//   1: index < 0: skip
//      index >= 0: [index, i) is 'type', status -> 0, index = -1
// ',':
//   0: skip
//   1: index >= 0: [index, i) is 'type', status -> 0, index = -1 (simulate 'str' is valid, index will >= 0 here)
// other chars:
//   0: index < 0: index = i
//      index >= 0: skip
//   1: index < 0: index = i
//      index >= 0: skip
func scan(str string) []*token {
	str += " " // for last 'type', todo: if any program skill to handle this situation?

	tokenList := make([]*token, 0, strings.Count(str, ",")+1)
	status := ScanStatus_ReadingName
	index := -1
	for i := 0; i < len(str); i++ {
		switch str[i] {
		case ' ':
			if index < 0 {
				continue
			}

			tokenItem := &token{
				value: str[index:i],
			}

			switch status {
			case ScanStatus_ReadingName:
				tokenItem.typ = TokenType_VariableName
				status = ScanStatus_ReadingType
			case ScanStatus_ReadingType:
				tokenItem.typ = TokenType_VariableType
				status = ScanStatus_ReadingName
			}

			tokenList = append(tokenList, tokenItem)

			index = -1
		case ',':
			if status != ScanStatus_ReadingType || index < 0 {
				continue
			}

			tokenList = append(tokenList, &token{
				value: str[index:i],
				typ:   TokenType_VariableType,
			})

			status = ScanStatus_ReadingName

			index = -1
		default:
			if index < 0 {
				index = i
			}
		}
	}

	return tokenList
}

func parseTokens(tokens []*token) []*param {
	if len(tokens)%2 != 0 {
		return nil
	}

	paramList := make([]*param, 0, len(tokens)/2)
	for i := 0; i+1 < len(tokens); i += 2 {
		paramItem := &param{
			name: tokens[i].value,
		}

		variableType := tokens[i+1].value
		if strings.HasPrefix(variableType, "[") {
			index := strings.LastIndex(variableType, "]")

			paramItem.isArray = true
			paramItem.arrayFlag = variableType[:index+1]

			variableType = variableType[index+1:]
		}
		if strings.HasPrefix(variableType, "*") {
			paramItem.isPointer = true

			variableType = variableType[1:]
		}

		paramItem.baseType = variableType

		paramList = append(paramList, paramItem)
	}

	return paramList
}

// string return the param like '[name] [type]'
func (p *param) string() string {
	res := []byte(p.name + " ")

	res = append(res, p.typ()...)

	return string(res)
}

// typ return the type of param
func (p *param) typ() string {
	var res []byte

	if p.isArray {
		res = append(res, p.arrayFlag...)
	}
	if p.isPointer {
		res = append(res, '*')
	}

	res = append(res, p.baseType...)

	return string(res)
}
