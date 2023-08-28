package main

import (
	"fmt"
	"reflect"
	"testing"
)

type Student struct {
	Name string "学生姓名"
	Age  int    `a:"1111" b:"3333"`
}

func TestReflectTypeOf(t *testing.T) {
	s := Student{}
	rt := reflect.TypeOf(s)
	fileName, ok := rt.FieldByName("Name")
	if ok {
		fmt.Println(fileName.Tag)
	}
	fieldAge, ok2 := rt.FieldByName("Age")
	if ok2 {
		fmt.Println(fieldAge.Tag.Get("a"))
		fmt.Println(fieldAge.Tag.Get("b"))
	}

	fmt.Println("type_name:", rt.Name())
	fmt.Println("type_numField:", rt.NumField())
	fmt.Println("type_PkgPath:", rt.PkgPath())
	println(rt.String())

	s.Age = 10
	s.Name = "Gavin"
}

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) String() {
	println("User:", u.Id, u.Name, u.Age)
}

func Info(o interface{}) {
	v := reflect.ValueOf(o)
	t := v.Type()
	println("Type:", t.Name())

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i).Interface()
		switch value := value.(type) {
		case int:
			fmt.Printf("%6s: %v = %d\n", field.Name, field.Type, value)
		case string:
			fmt.Printf("%6s: %v = %s", field.Name, field.Type, value)
		}
	}
}

func TestReflectValue(t *testing.T) {
	u := User{1, "Tome", 30}
	Info(u)
}
