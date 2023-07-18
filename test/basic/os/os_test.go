package main

import (
	"fmt"
	"os"
	"testing"
)

func TestOsStat(t *testing.T) {
	fileInfo, _ := os.Stat("/Users/songbowen/go/src/learning-go")
	println(fileInfo.Name())
	println(fileInfo.IsDir())
	if _, err := os.Stat("hello"); err != nil {
		fmt.Println("file hello not exits")
		fmt.Println(err)
	}
}
