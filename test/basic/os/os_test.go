package main

import (
	"fmt"
	"os"
	"testing"
)

func TestGetWd(t *testing.T) {
	dir, _ := os.Getwd()
	fmt.Println(dir)
}

func TestOsStat(t *testing.T) {
	fileInfo, _ := os.Stat("/Users/jiayun/.kube/sit-config")
	println(fileInfo.Name())
	println(fileInfo.IsDir())
	if _, err := os.Stat("hello"); err != nil {
		fmt.Println("file hello not exits")
		fmt.Println(err)
	}
}

func TestOsExecutable(t *testing.T) {
	executable, err := os.Executable()
	if err != nil {
		return
	}
	println(fmt.Sprintf("打印编译成的二进制文件的绝对路径=%v", executable))
}
