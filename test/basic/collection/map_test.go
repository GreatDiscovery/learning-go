package main

import (
	"fmt"
	"testing"
)

type person struct {
	Name string
	Age  int
}

func (p person) string() string {
	return fmt.Sprintf("name=%s, age=%v", p.Name, p.Age)
}

func TestMapPointer(t *testing.T) {
	m := make(map[string]*person)
	if v, ok := m["k1"]; ok {
		fmt.Println("ok, v", v)
	} else {
		fmt.Println("not ok, v", v)
	}
}
