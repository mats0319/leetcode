package main

import (
    "bytes"
    "flag"
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
    InputParams string

    FuncNameExported string
    FuncNameActually string
    FuncNameSpelled string
}

const (
    fileName = "main_test.go"
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
			t.Errorf("{{.FuncNameSpelled}} test failed on case: %d", i)
		}
	}
}
`
)

var (
    help            bool
    funcDeclaration string
    questionNumber int
)

func init() {
    flag.BoolVar(&help, "h", false, "this help")
    flag.StringVar(&funcDeclaration, "d", "", "give a function declaration,\n" +
        "with template '{func name}({input type[s]}){output type}'\n" +
        "e.g. 'testGen(int,string)[]byte'")
    flag.IntVar(&questionNumber, "n", -1, "give out the question number to redirect the generated file\n" +
        "default path is the same as this tool\n" +
        "please generate path like 'Solutions' dir in this repo before.")

    flag.Parse()

    if help {
        log.Println("Options: ")
        flag.PrintDefaults()
        os.Exit(0)
    }

    if funcDeclaration == "" {
        log.Fatalln("Please input a function declaration. ")
    }
}

func main() {
    log.Println("test file generate tool started. ")

    generateFile(replaceTemplate(formatFuncDeclaration()))

    log.Println("test file generate tool finished. ")
}

func formatFuncDeclaration() (result *FunctionDeclaration) {
    formatDeclaration := trimAllSpaceAndSplit() // fd: func name + inputs + output
    result = &FunctionDeclaration{
        ExpectType:       formatDeclaration[2],
        FuncNameActually: formatDeclaration[0],
    }

    result.formatInputs(formatDeclaration[1])
    result.formatFuncName(formatDeclaration[0])

    return
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
                fd.InputTypes = "[]"+fd.InputTypes
                for i := 0; i < m[key]; i++ {
                    fd.InputParams += ", tcs[i].In["+strconv.Itoa(i)+"]"
                }
                fd.InputParams = fd.InputParams[2:]
            }
        }
    } else { // len(m) > 1
        fd.InputTypes = "In"

        fd.InputStructDefine = "type In struct {\n"
        for key := range m {
            for i := 0; i < m[key]; i++ {
                fd.InputStructDefine += string(key[0]-'a'+'A')+key[1:]+" "+key+"\n"
            }
        }
        fd.InputStructDefine += "}\n"

        for i := range inputsSplit {
            fd.InputParams += ", tcs[i].In."+string(inputsSplit[i][0]-'a'+'A')+inputsSplit[i][1:]
        }
        fd.InputParams = fd.InputParams[2:]
    }

    return
}

func (fd *FunctionDeclaration) formatFuncName(funcName string) {
    if 'a' <= funcName[0] && funcName[0] <= 'z' {
        funcName = string(funcName[0]+'A'-'a')+funcName[1:]
    }
    fd.FuncNameExported = funcName

    var (
        funcNameSplit = make([]byte, len(funcName)*2)
        t int
    )
    for i := 0; i < len(funcName); i, t = i+1, t+1 {
        if 'A' <= funcName[i] && funcName[i] <= 'Z' {
            funcNameSplit[t] = ' '
            t++
            funcNameSplit[t] = funcName[i]-'A'+'a'
        } else {
            funcNameSplit[t] = funcName[i]
        }
    }
    fd.FuncNameSpelled = string(funcNameSplit[1:t])

    return
}

func trimAllSpaceAndSplit() (result []string) {
    for i := 0; i < len(funcDeclaration); i++ {
        if funcDeclaration[i] == ' ' {
            funcDeclaration = funcDeclaration[:i] + funcDeclaration[i+1:]
        }
    }

    var t int
    for i := range funcDeclaration {
        if funcDeclaration[i] == '(' {
            t = i
            result = append(result, funcDeclaration[:i])
            break
        }
    }
    for i := t+1; i < len(funcDeclaration); i++ {
        if funcDeclaration[i] == ')' {
            result = append(result, funcDeclaration[t+1: i])
            result = append(result, funcDeclaration[i+1:])
            break
        }
    }

    return
}

func replaceTemplate(fd *FunctionDeclaration) (result []byte) {
    t, err := template.New("test genarate").Parse(templateStr)
    if err != nil {
        log.Fatalln(err)
    }

    buff := bytes.NewBufferString("")
    if err = t.Execute(buff, fd); err != nil {
        log.Fatalln(err)
    }

    result, err = format.Source(buff.Bytes())
    if err != nil {
        log.Fatalln(err)
    }

    return
}

func generateFile(src []byte) {
    var dir string
    if questionNumber != -1 {
        dir = generateDir()
    } else {
        dir, _ = os.Getwd()
    }
    fullName := dir +"\\" + fileName
    if _, err := os.Stat(fullName); os.IsNotExist(err) {
        if err = ioutil.WriteFile(fullName, src, 0755); err != nil {
            log.Fatalln("write failed", err)
        }
    } else {
        log.Fatalln("file exist", fileName)
    }

    return
}

func generateDir() string {

    return ""
}
