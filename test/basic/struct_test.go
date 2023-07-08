package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Books struct {
	Title   string `json:"title,omitempty"`
	Author  string `json:"author,omitempty"`
	Subject string `json:"subject,omitempty"`
	BookId  int    `json:"book_id,omitempty"`
}

func TestStruct(t *testing.T) {

	// 创建一个新的结构体
	fmt.Println(Books{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407})

	// 也可以使用 key => value 格式
	fmt.Println(Books{Title: "Go 语言", Author: "www.runoob.com", Subject: "Go 语言教程", BookId: 6495407})

	// 忽略的字段为 0 或 空
	fmt.Println(Books{Title: "Go 语言", Author: "www.runoob.com"})
}

func TestJson(t *testing.T) {
	book := Books{
		Title:   "钢铁是怎样练成的",
		Author:  "奥斯特洛夫斯基",
		Subject: "novel",
		BookId:  1,
	}

	data, err := json.Marshal(book)
	if err != nil {
		return
	}
	fmt.Println(string(data))
}
