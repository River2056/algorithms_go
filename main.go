package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	name := flag.String("name", "", "file path")
	flag.Parse()

	if *name == "" {
		fmt.Println("Please enter file path!")
		return
	}

	arr := strings.Split(*name, "/")
	path := filepath.Join(arr...)
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		panic(err)
	}

	tmp := make([]string, len(arr))
	copy(tmp, arr)
	tmp = append(tmp, "main.go")
	mainFile := filepath.Join(tmp...)

	f, err := os.Create(mainFile)
	if err != nil {
		panic(err)
	}
	f.Close()

	tmp = make([]string, len(arr))
	copy(tmp, arr)
	testFile := path[strings.LastIndex(path, "/")+1:]
    testFile = testFile + "_test.go"
	tmp = append(tmp, testFile)
    testFile = filepath.Join(tmp...)

    f, err = os.Create(testFile)
    if err != nil {
        panic(err)
    }
    f.Close()
}
