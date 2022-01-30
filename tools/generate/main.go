// Test_gen can generate universal leetcode test file,
// input a question's number with flag '-n', and we will generate the test file and redirect to it's position.
//
// Once more, We support universal questions only, some special questions are not support, for example:
// - return value's type can't use '==' to judge equal
// - param type is double-level slice
//
// Mario, 2019.12.24
//
// Last valid modify: (valid modify: except re-name, format code and so on)
// 2020.1.27

package main

import (
	"flag"
	"log"
	"os"
)

var (
	help        bool
	questionNum int
)

func init() {
	flag.BoolVar(&help, "h", false, "this help")
	flag.IntVar(&questionNum, "n", -1, "give out question number\n"+
		"we will read function declaration and redirect generated file automatically\n"+
		"a valid number is from 0 to 9999 (0 means generate test at default place).")

	flag.Parse()

	if help {
		log.Println("Options: ")
		flag.PrintDefaults()
		os.Exit(0)
	}

	if questionNum < 0 || questionNum >= 10000 {
		log.Fatalln("Please input a valid question's number, more information in help with '-h' flag. ")
	}
}

func main() {
	log.Println("test file generate tool started. ")

	dir := getDir(questionNum)
	funcDeclaration := getFuncDeclaration(dir + fileName)
	formatDeclaration := formatFuncDeclaration(funcDeclaration)

	payload := replaceTemplate(formatDeclaration)

	saveFile(payload, dir+testFileName)

	log.Println("test file generate tool finished. ")
}
