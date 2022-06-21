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
// Last modify: 2022.6.20

package main

import (
	"flag"
	"log"
	"os"
)

var (
	help        bool
	questionNum int
	createDir   bool
)

func init() {
	flag.BoolVar(&help, "h", false, "this help")
	flag.IntVar(&questionNum, "n", -1, "question number\n"+
		"use for read function declaration and redirect generated file\n"+
		"a valid number is from 0 to 9999 (0 is test for this generator).")
	flag.BoolVar(&createDir, "c", false, "create dir structure only, according to question number.")

	flag.Parse()

	if help {
		log.Println("Options: ")
		flag.PrintDefaults()
		os.Exit(0)
	}

	if questionNum < 0 || questionNum >= 10000 || (createDir && questionNum < 1) {
		log.Fatalln("Please input a valid question number, more information in help with '-h' flag. ")
	}
}

type templateData struct {
	// params type in test case
	InputTypes string
	ExpectType string

	InputStructDefine string // if input params' type more than 1, define struct
	InputParams       string

	FuncExportedName string // part of test function name, like 'Testxxx', without 'Test-' prefix
	FuncActuallyName string // for invoke function
	FuncSpelledName  string // in error message
}

func main() {
	log.Println("test file generator started. ")
	defer log.Println("test file generator exited. ")

	dir := generateDir(questionNum)

	if createDir {
		err := createFile(dir, fileName)
		if err != nil {
			log.Println("create file failed, error: ", err)
		}

		return
	}

	_, err := os.Stat(dir + testFileName)
	if err == nil { // file is exist
		log.Println("file is exist: "+ dir+testFileName)
		return
	}

	funcDeclaration := getFuncDeclaration(dir + fileName) // [func name]([input param(s)]) [output param(s)]
	templateDataIns := generateTemplateData(funcDeclaration)

	payload, err := replaceTemplate(templateDataIns)
	if err != nil {
		log.Println("replace template failed, error: ", err)
		return
	}

	err = saveFile(dir+testFileName, payload)
	if err != nil {
		log.Println("save file failed, error: ", err)
		return
	}
}
