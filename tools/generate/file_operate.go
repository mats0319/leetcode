package main

import (
    "bufio"
    "bytes"
    "fmt"
    "github.com/pkg/errors"
    "go/format"
    "html/template"
    "log"
    "os"
)

func getDir(num int) string {
    if num == 0 {
        return "./default/"
    }

    var result string
    for i := 10; i < 10000; i *= 10 {
        var tag int
        if num%i == 0 {
            tag = -1
        }
        result = fmt.Sprintf("%d-%d/", (num/i+tag)*i+1, (num/i+tag)*i+i) + result
    }

    return fmt.Sprintf("../../solutions/%s%d/", result, num)
}

func getFuncDeclaration(fileName string) string {
    file, err := os.Open(fileName)
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
            log.Fatalln("file scan finished, don't have a function to test") // file exit
        }

        // test first func
        if len(fileScanner.Text()) > 5 && fileScanner.Text()[:5] == "func " { // todo: RE replace
            break
        }
    }

    return fileScanner.Text()[5 : len(fileScanner.Text())-1]
}

func replaceTemplate(fd *FunctionDeclaration) []byte {
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

    return result
}

func saveFile(src []byte, fileName string) {
    err := os.WriteFile(fileName, src, 0755)
    if err != nil {
        log.Fatalln(errors.Wrap(err, "write failed"))
    }
}
