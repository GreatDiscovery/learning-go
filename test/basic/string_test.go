package main

import (
	"encoding/json"
	"fmt"
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

func TestJsonTest3(t *testing.T) {
	var a A
	str := "{\n    \"name\": \"\",\n    \"age\": 3\n} "
	err := json.Unmarshal([]byte(str), &a)
	//if err != nil {
	//	fmt.Printf("%+v\n", errors.Wrap(err, "打印堆栈"))
	//}
	if err != nil {
		println(err.Error())
	}
}
