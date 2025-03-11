package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

// 示例结构体
type Person struct {
	Name  string
	Age   int
	Email string
}

func TestReflect(t *testing.T) {
	// 创建一个结构体实例
	p := Person{}

	// 获取结构体的反射值
	value := reflect.ValueOf(&p).Elem() // 必须使用指针的 Elem() 获取可设置的反射值

	// 获取结构体的类型信息
	typ := value.Type()

	// 遍历结构体的字段
	for i := 0; i < value.NumField(); i++ {
		// 获取字段的名称和类型
		field := typ.Field(i)
		fieldName := field.Name
		fieldType := field.Type

		fmt.Printf("Field %d: Name = %s, Type = %s\n", i+1, fieldName, fieldType)
	}

	// 根据字段名称给结构体赋值
	setFieldByName(&p, "Name", "Alice")
	setFieldByName(&p, "Age", 30)
	setFieldByName(&p, "Email", "alice@example.com")

	// 打印赋值后的结构体
	fmt.Printf("Updated Person: %+v\n", p)
	assert.Equal(t, "Alice", p.Name)
	assert.Equal(t, 30, p.Age)
	assert.Equal(t, "alice@example.com", p.Email)
}

// 根据字段名称给结构体赋值
func setFieldByName(obj interface{}, fieldName string, value interface{}) {
	// 获取结构体的反射值
	v := reflect.ValueOf(obj).Elem()

	// 获取字段的反射值
	field := v.FieldByName(fieldName)

	// 检查字段是否存在
	if !field.IsValid() {
		fmt.Printf("Field %s not found\n", fieldName)
		return
	}

	// 检查字段是否可设置
	if !field.CanSet() {
		fmt.Printf("Field %s cannot be set\n", fieldName)
		return
	}

	// 获取值的反射值
	val := reflect.ValueOf(value)

	// 检查类型是否匹配
	if field.Type() != val.Type() {
		fmt.Printf("Type mismatch: expected %s, got %s\n", field.Type(), val.Type())
		return
	}

	// 给字段赋值
	field.Set(val)
}
