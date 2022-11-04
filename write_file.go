package main

import (
	"fmt"
	"os"
)

func writeInFile(fileName string, content []byte) (err error) {

	f, err := os.Create(fileName)

	if err != nil {
		fp, n, l := getFilePathFuncNameAndLine()
		err = fmt.Errorf(fp, n, l, err)
		return
	}
	defer f.Close()

	_, err = f.Write(content)
	return err
}
