package main

import (
	"fmt"
	"testing"
)

type Day int

const (
	Sunday Day = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

// 定义 Day 类型的 String 方法
func (d Day) String() string {
	names := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	if d < Sunday || d > Saturday {
		return "Unknown"
	}
	return names[d]
}

func TestEnum(t *testing.T) {
	fmt.Println(Sunday)    // 输出: Sunday
	fmt.Println(Wednesday) // 输出: Wednesday
	fmt.Println(Day(7))    // 输出: Unknown
}
