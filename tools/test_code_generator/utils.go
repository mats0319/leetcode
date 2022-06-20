package main

import (
    "fmt"
    "github.com/pkg/errors"
    "log"
    "os"
)

func generateDir(num int) string {
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

// createFile create file if not exist
func createFile(dir string, fileName string) error {
    err := os.MkdirAll(dir, 0755)
    if err != nil {
        log.Println("mkdir failed, error: ", err)
        return err
    }

    _, err = os.Stat(dir + fileName)
    if err == nil { // file is exist
        log.Println("file is exist: "+ dir+fileName)
        return errors.New("file is exist: "+ dir+fileName)
    }

    _, err = os.Create(dir + fileName)
    if err != nil {
        log.Println("create file failed, error: ", err)
        return err
    }

    return nil
}

func trimAllSpace(str string) string {
    byteSlice := make([]byte, 0, len(str))
    for i := 0; i < len(str); i++ {
        if str[i] != ' ' {
            byteSlice = append(byteSlice, str[i])
        }
    }

    return string(byteSlice)
}

func makeFirstCharBigCase(str string) string {
    byteSlice := []byte(str)

    if 'a' <= byteSlice[0] && byteSlice[0] <= 'z' {
        byteSlice[0] = byteSlice[0]-'a'+'A'
    }

    return string(byteSlice)
}
