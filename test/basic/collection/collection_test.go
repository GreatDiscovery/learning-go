package main

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	var x = make([]int, 3, 5)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
