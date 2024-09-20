package main

import (
	"fmt"
	"strings"
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

func DayOf(day string) Day {
	switch strings.ToLower(day) {
	case "sunday":
		return Sunday
	case "monday":
		return Monday
	case "tuesday":
		return Tuesday
	case "wednesday":
		return Wednesday
	case "thursday":
		return Thursday
	case "friday":
		return Friday
	case "saturday":
		return Saturday
	}
	return Sunday
}

func TestEnum(t *testing.T) {
	fmt.Println(Sunday)    // 输出: Sunday
	fmt.Println(Wednesday) // 输出: Wednesday
	fmt.Println(Day(7))    // 输出: Unknown
	fmt.Println(DayOf("friday"))
}
