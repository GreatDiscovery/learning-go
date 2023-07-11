package main

import (
	"fmt"
	"testing"
)

// define interface and implementation

type Phone interface {
	call()
}

type Vivo struct{}
type Iphone struct{}

func (vivo Vivo) call() {
	fmt.Println("vivo call phone!")
}

func (apple Iphone) call() {
	fmt.Println("apple iphone call!")
}

func TestInterface(t *testing.T) {
	var phone Phone
	phone = new(Vivo)
	phone.call()

	phone = new(Iphone)
	phone.call()
}

//  type assertion，接口断言，就是判断接口是否 是某种类型

type Inter interface {
	Ping()
	Pang()
}

type Anter interface {
	Inter
	String()
}

type St struct {
	Name string
}

func (St) Ping() {
	println("ping")
}

func (St) Pang() {
	println("pang")
}

func TestTypeAssert(t *testing.T) {
	st := &St{"andes"}
	var i interface{} = st

	// 判断i绑定的实例是否实现了接口类型Inter
	o := i.(Inter)
	o.Ping()
	o.Pang()

	// i没有实现接口Anter，会报错
	//p := i.(Anter)
	//p.String()

	// 判断i实现了具体类型St
	s := i.(*St)
	fmt.Printf("%s", s.Name)

}
