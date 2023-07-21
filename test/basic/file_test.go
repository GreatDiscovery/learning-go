package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/go-ini/ini"
	"os"
	"testing"
)

var (
	conf = flag.String("conf", "../../conf/dev.ini", "conf")
)

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
