package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/go-ini/ini"
	"os"
	"path/filepath"
	"testing"
)

var (
	conf = flag.String("conf", "../../conf/dev.ini", "conf")
)

func TestPwd(t *testing.T) {
	abs, err := filepath.Abs(".")
	if err != nil {
	}
	fmt.Println("pwd is ")
	println(abs)
}

func TestCreateFile(t *testing.T) {
	if err := CreateDirAndFile(); err != nil {
		panic(err)
	}
}

func CreateDirAndFile() (err error) {
	var paths []string
	defer func() {
		if err != nil {
			for _, d := range paths {
				os.RemoveAll(d)
			}
		}
	}()
	root := "/Users/songbowen/Documents/github/learning-go/tmp"
	dir1 := filepath.Join(root, "dir1")

	paths = append(paths, dir1)
	if err := os.MkdirAll(dir1, 0777); err != nil {
		return err
	}

	file1 := filepath.Join(dir1, "file1")
	paths = append(paths, file1)
	content := "write to file1, hello world! "
	err = os.WriteFile(file1, []byte(content), 0777)
	if err != nil {
		return err
	}

	contentBytes, err := os.ReadFile(file1)
	if err != nil {
		return err
	}
	println(fmt.Sprintf("file content is %v", string(contentBytes)))

	for _, d := range paths {
		os.RemoveAll(d)
	}
	return nil
}

func TestIni(t *testing.T) {
	flag.Parse()
	cfg, err := ini.Load(*conf)
	if err != nil {
		println(err.Error())
	}
	println(cfg.SectionStrings())
}

func TestWriteFile(t *testing.T) {
	//创建一个新文件，写入内容 5 句 “http://c.biancheng.net/golang/”
	filePath := "/Users/jiayun/Downloads/tmp/1.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	//及时关闭file句柄
	defer file.Close()
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		write.WriteString("http://c.biancheng.net/golang/\n")
	}
	//Flush将缓存的文件真正写入到文件中
	write.Flush()
}
