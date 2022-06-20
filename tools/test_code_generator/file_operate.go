package main

import (
	"bufio"
	"bytes"
	"github.com/pkg/errors"
	"go/format"
	"html/template"
	"log"
	"os"
)

func getFuncDeclaration(fileName string) string {
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

		// test first func
		if len(fileScanner.Text()) > 5 && fileScanner.Text()[:5] == "func " { // todo: RE replace
			break
		}
	}

	return fileScanner.Text()[5 : len(fileScanner.Text())-1]
}

func replaceTemplate(data *templateData) ([]byte, error) {
	t, err := template.New("test file generate tool").Parse(templateStr)
	if err != nil {
		log.Println("parse template failed, error: ", err)
		return nil, err
	}

	buff := bytes.NewBufferString("")
	if err = t.Execute(buff, data); err != nil {
		log.Println("fill template failed, error: ", err)
		return nil, err
	}

	result, err := format.Source(buff.Bytes())
	if err != nil {
		log.Println("format source failed, error: ", err)
		return nil, err
	}

	return result, nil
}

func saveFile(fileName string, src []byte) error {
	err := os.WriteFile(fileName, src, 0755)
	if err != nil {
		log.Println(errors.Wrap(err, "write failed"))
		return err
	}

	return nil
}
