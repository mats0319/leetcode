// Test Code Generator can generate test file of universal leetcode questions
//
// Author: Mario
//
// We support universal situations only, some special situations required manual modification, for example:
// - design questions, different format of test code
// - return value type not define on '==': write compare function on return value type
//
// tips:
// 1. We consider return value list only contains one item and without name,
//    because leetcode support many programming languages and some of them not support multi-value return or named
//    return value, so leetcode has no reason to define go function with these special go-style return values
//    further, we do not support them
// 2. We ignore the situation that multi input params use one type declaration, like 'num1, num2 int',
//    reasons are same as above
//
// First version: 2019.12.24
// Last modify: 2022.11.4

package main

import (
	"flag"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"os"
)

var (
	help              bool
	questionNum       int
	initPath          bool
	overwriteTestFile bool
)

func init() {
	flag.BoolVar(&help, "h", false, "this help")
	flag.IntVar(&questionNum, "n", -1, "question number\n"+
		"use for read function declaration and redirect generated file\n"+
		"a valid number is from 0 to 9999 (0 is test for this generator)")
	flag.BoolVar(&initPath, "i", false, "initialize dir structure and 'main.go' file, according to 'question number'")
	flag.BoolVar(&overwriteTestFile, "f", false, "force generate test file, overwrite if it exist")

	flag.Parse()

	if help {
		log.Println("Options: ")
		flag.PrintDefaults()
		os.Exit(0)
	}

	if questionNum < 0 || questionNum >= 10000 || (initPath && questionNum < 1) {
		log.Fatalln("Please input a valid question number, more information in help with '-h' flag. ")
	}
}

func main() {
	log.Println("test file generator started. ")
	defer log.Println("test file generator exited. ")

	dir := calcDir(questionNum)

	if initPath {
		err := createFile(dir, fileName)
		if err != nil {
			log.Println("create file failed, error: ", err)
		}

		return
	}

	if !overwriteTestFile {
		_, err := os.Stat(dir + testFileName)
		if err == nil { // file is exist
			log.Println("file is exist: " + dir + testFileName)
			return
		}
	}

	templateDataIns := generateTemplateData(parseFuncDeclaration(dir + fileName))

	payload, err := templateDataIns.fillTemplate()
	if err != nil {
		log.Println("replace template failed, error: ", err)
		return
	}

	err = os.WriteFile(dir+testFileName, payload, 0755)
	if err != nil {
		log.Println("save file failed, error: ", errors.Wrap(err, ""))
		return
	}

	return
}

func calcDir(num int) string {
	if num == 0 {
		return "./testdata/"
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
		log.Println("file is exist: " + dir + fileName)
		return errors.New("file is exist: " + dir + fileName)
	}

	_, err = os.Create(dir + fileName)
	if err != nil {
		log.Println("create file failed, error: ", err)
		return err
	}

	return nil
}
