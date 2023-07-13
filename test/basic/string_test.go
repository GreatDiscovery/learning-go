package main

import (
	"encoding/json"
	"fmt"
	"github.com/json-iterator/go"
	"strings"
	"testing"
)

type Body struct {
	Code    int                    `json:"code"`
	Data    map[string]interface{} `json:"data,omitempty"`
	Message string                 `json:"message"`
	Success int                    `json:"success"`
}

// 去除json转换过程中的换行符
func TestRemoveLF(t *testing.T) {
	var body Body
	str := "{\n  \"code\": -400,\n  \"data\": \"\",\n  \"message\": \"NodeController.node_type error\",\n  \"success\": 0\n}"
	b := []byte(str)
	err := json.Unmarshal(b, &body)
	if err != nil {
		println(err.Error())
	}
	replace := strings.Replace(str, "\n", "", -1)
	fmt.Printf("replace=%s", replace)
	fmt.Println()
	b = []byte(replace)
	err = json.Unmarshal(b, &body)
	if err != nil {
		println(err.Error())
	} else {
		println(body.Message)
	}
}

func TestRemoveLF2(t *testing.T) {
	var body Body
	str := "{\n    \"code\": -400,\n    \"data\": \"\",\n    \"message\": \"NodeController.node_type error\",\n    \"success\": 0\n}"
	b := []byte(str)
	err := json.Unmarshal(b, &body)
	if err != nil {
		println(err.Error())
	} else {
		println(body.Message)
	}
}

type A struct {
	Name map[string]interface{} `json:"name"`
	Age  int64                  `json:"age"`
}

func TestJsoner(t *testing.T) {
	var a A
	str := "{\n    \"name\": \"\",\n    \"age\": 3\n}"
	err := jsoniter.Unmarshal([]byte(str), &a)
	if err != nil {
		println(err.Error())
	} else {
		println("success")
	}
}

func TestJson3(t *testing.T) {
	var a A
	str := "{\n    \"name\": \"\",\n    \"age\": 3\n}"
	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		println(err.Error())
	}
}

type B struct {
	Data interface{}
	Code int64
}

func TestTypeAssert2(t *testing.T) {
	var b B
	str := "{\n    \"Data\": \"\",\n    \"Code\": 200\n} "
	err := json.Unmarshal([]byte(str), &b)
	if err != nil {
		println(err.Error())
	}
	println(typeof(b.Data))

	println(typeof(b.Code))

	var b2 B
	str2 := "{\n    \"Data\": {\n        \"name\": \"jiayun\",\n        \"age\": 18\n    },\n    \"Code\": 200\n} "
	err2 := json.Unmarshal([]byte(str2), &b2)
	if err2 != nil {
		println(err2.Error())
	}
	println(typeof(b2.Data))
	println(typeof(b2.Code))

	switch b2.Data.(type) {
	case string:
		println("type is string")
	case map[string]interface{}:
		println("type is map")
	}

}

func typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}