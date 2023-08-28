package main

import (
	"fmt"
	"testing"
)

// for range是对象的拷贝，不能直接修改元对象。需要使用数组指针修改源对象
func TestForModify(t *testing.T) {
	b1 := Books{
		Title:   "1",
		Author:  "jiayun",
		Subject: "science",
		BookId:  1,
	}
	b2 := Books{
		Title:   "2",
		Author:  "jiayun",
		Subject: "computer",
		BookId:  2,
	}
	var books []Books
	books = append(books, b1)
	books = append(books, b2)

	for _, book := range books {
		if book.BookId == 1 {
			book.BookId = 3
		}
	}
	fmt.Printf("books=%v\n", books)

	for i, book := range books {
		if book.BookId == 1 {
			books[i].BookId = 3
		}
	}
	fmt.Printf("books=%v", books)
}
