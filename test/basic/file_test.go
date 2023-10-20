package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/go-ini/ini"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"path/filepath"
	"syscall"
	"testing"
)

var (
	conf = flag.String("conf", "../../conf/dev.ini", "conf")
)

func TestFileUtils(t *testing.T) {
	fileName := "/Users/jiayun/Documents/github/learning-go/test/framework/emptydir/tmp.1"
	//defer os.Remove(fileName)
	context := "hello world"
	var d = []byte(context)
	// you must make dir before use this writeFile function
	err := ioutil.WriteFile(fileName, d, 666)
	if err != nil {
		fmt.Println(err)
	}
}

func TestFileExist(t *testing.T) {
	fileName1 := "tmp.1"
	fileName2 := "tmp.2"
	defer os.Remove(fileName1)
	defer os.Remove(fileName2)

	f1, err := os.OpenFile(fileName1, syscall.O_WRONLY|syscall.O_CREAT|syscall.O_TRUNC, 0666)
	if err != nil {
		return
	}
	f1.Sync()
	var f1Exist bool
	var f2Exist bool
	_, err = os.Stat(fileName1)
	if err != nil {
		f1Exist = false
	} else {
		f1Exist = true
	}
	_, err = os.Stat(fileName2)
	if err != nil {
		f2Exist = false
	} else {
		f2Exist = true
	}

	assert.Equal(t, true, f1Exist)
	assert.Equal(t, false, f2Exist)

}

func TestFileMode(t *testing.T) {
	fileName := "tmp.1"
	defer os.Remove(fileName)
	context := "hello world"
	var d = []byte(context)
	// 0666表示文件的权限
	err := os.WriteFile(fileName, d, 0666)
	if err != nil {
		log.Error("write fail: " + err.Error())
		return
	}
}

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
