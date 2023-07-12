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

func (b Books) String() string {
	marshal, err := json.Marshal(b)
	if err != nil {
		return ""
	}
	return string(marshal)
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
	fmt.Println("------------")
	fmt.Println(book.String())
}

type LowerCase struct {
	title  string
	author string
}

func TestStructLowerCase(t *testing.T) {
	book := LowerCase{
		title:  "平凡的世界",
		author: "路遥",
	}

	data, err := json.Marshal(book)
	if err != nil {
		return
	}
	// 首字母小写无法json解析，因为首字母小写权限是私有的
	println(string(data))
}

type Base struct {
}

// @Deprecated 该方法由于不知道子类的类型，所以无法解析
func (receiver *Base) JsonString() string {
	data, err := json.Marshal(receiver)
	if err != nil {
		return ""
	}
	return string(data)
}

type BaseA struct {
	Base
	Name string `json:"name"`
}

type BaseB struct {
	Base
	Name string `json:"name"`
}

func JsonString(str interface{}) string {
	data, err := json.Marshal(str)
	if err != nil {
		return ""
	}
	return string(data)
}

func TestFuncInherit(t *testing.T) {
	a := BaseA{Name: "BaseA"}
	b := BaseB{Name: "BaseB"}
	println(JsonString(a))
	println(JsonString(b))
}
