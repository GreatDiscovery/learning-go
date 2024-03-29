package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/assert/v2"
	"sort"
	"testing"
)

type Books struct {
	Title   string `json:"title,omitempty"`
	Author  string `json:"author,omitempty"`
	Subject string `json:"subject,omitempty"`
	BookId  int    `json:"book_id,omitempty"`
}

func (b Books) Clone() *Books {
	copy := Books{
		Title:   b.Title,
		Author:  b.Author,
		Subject: b.Subject,
		BookId:  b.BookId,
	}
	return &copy
}

func TestClone(t *testing.T) {
	b1 := Books{
		Title:   "钢铁是怎样练成的",
		Author:  "奥斯特洛夫斯基",
		Subject: "novel",
		BookId:  1,
	}
	b2 := b1.Clone()
	b2.Author = "李逵"
	assert.NotEqual(t, b1, b2)
}

// sort1
type BookSlice []Books

func (b BookSlice) Len() int {
	return len(b)
}
func (b BookSlice) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
func (b BookSlice) Less(i, j int) bool { // 重写 Less() 方法， 从大到小排序
	return b[i].BookId < b[j].BookId
}

func (b Books) String() string {
	marshal, err := json.Marshal(b)
	if err != nil {
		return ""
	}
	return string(marshal)
}

func TestSortStruct1(t *testing.T) {
	books := []Books{
		{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407},
		{
			Title:   "钢铁是怎样练成的",
			Author:  "奥斯特洛夫斯基",
			Subject: "novel",
			BookId:  1,
		},
		{
			Title:  "平凡的世界",
			Author: "路遥",
			BookId: 3,
		},
	}
	// sort1
	sort.Sort(BookSlice(books))
	fmt.Println(books)

	// sort2
	sort.Slice(books, func(i, j int) bool {
		return books[i].Title < books[j].Title
	})
	fmt.Println(books)

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
