package main

import (
    "fmt"
    "strconv"
    "strings"
)

type FunctionDeclaration struct {
    InputTypes string
    ExpectType string

    InputStructDefine string // if input params' type more than 1
    InputParams       string

    FuncNameExported string // for 'Testxxx' func name
    FuncNameActually string // for call func
    FuncNameSpelled  string // for error message
}

func formatFuncDeclaration(funcDeclaration string) *FunctionDeclaration {
    formatDeclaration := splitFuncDeclaration(funcDeclaration) // fd = func name + inputs + output
    result := &FunctionDeclaration{
        ExpectType:       formatDeclaration[2],
        FuncNameActually: formatDeclaration[0],
    }

    result.formatInputs(formatDeclaration[1])
    result.formatFuncNames(formatDeclaration[0])

    return result
}

func splitFuncDeclaration(funcDeclaration string) (result []string) {
    var t int
    for t < len(funcDeclaration) && funcDeclaration[t] != '(' {
        t++
    }
    result = append(result, funcDeclaration[:t]) // func name

    for i := t + 1; i < len(funcDeclaration); i++ {
        if funcDeclaration[i] == ')' {
            result = append(result, funcDeclaration[t+1:i]) // input param(s)
            result = append(result, funcDeclaration[i+1:])  // output param(s)
            break
        }
    }

    result[1] = removeVariableName(result[1])

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
                    fd.InputStructDefine = fmt.Sprintf("%s%sSlice%d %s\n", fd.InputStructDefine, typeToField(key), i, key)
                }
            }
        }
        fd.InputStructDefine = fmt.Sprintf("type In struct {\n%s}\n", fd.InputStructDefine)

        for i := range inputsSplit {
            if 'a' <= inputsSplit[i][0] && inputsSplit[i][0] <= 'z' {
                fd.InputParams = fmt.Sprintf("%s, tcs[i].In.%s%s%d", fd.InputParams, string(inputsSplit[i][0]-'a'+'A'), inputsSplit[i][1:], m[inputsSplit[i]]-1)
            } else { // param is a slice
                fd.InputParams = fmt.Sprintf("%s, tcs[i].In.%sSlice%d", fd.InputParams, typeToField(inputsSplit[i]), m[inputsSplit[i]]-1)
            }
            m[inputsSplit[i]]--
        }
        fd.InputParams = fd.InputParams[2:]
    }

    return
}

func typeToField(str string) (field string) {
    for i, c := range str {
        if c == '[' || c == ']' {
            continue
        }
        bigCase := str[i] - 'a' + 'A'
        field = string(bigCase) + str[i+1:]
        break
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
