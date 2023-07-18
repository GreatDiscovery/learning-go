package main

import (
	"fmt"
	"testing"
)

// 区分方法和函数
// 函数如下
func add(a, b int) int {
	return a + b
}

func TestFunc(t *testing.T) {
	sum := add(1, 2)
	fmt.Print(sum)
}

// 方法有明确的接收者，这样就把接收者和方法绑定到了一起
type person struct {
	name string
}

func (p person) String() string {
	return "name is " + p.name
}

func TestMethod(t *testing.T) {
	p := person{"gavin"}
	fmt.Print(p.String())
}

// 区别值传递还是指针传递
func (p person) modifyNameValue(name string) {
	p.name = name
}

func (p *person) modifyNameReference(name string) {
	p.name = name
}

func TestModify(t *testing.T) {
	p := person{
		"gavin",
	}

	p.modifyNameValue("bob")
	fmt.Println(p.String())

	p.modifyNameReference("bob")
	fmt.Print(p.String())
}
