package main

import (
	"encoding/json"
	"fmt"
	"github.com/json-iterator/go"
	"strconv"
	"strings"
	"testing"
)

type Body struct {
	Code    int                    `json:"code"`
	Data    map[string]interface{} `json:"data,omitempty"`
	Message string                 `json:"message"`
	Success int                    `json:"success"`
}

// 测试string和byte[]之间互相转换
func TestStr2Byte(t *testing.T) {
	str1 := "hello world!"
	arr1 := []byte(str1)
	str2 := string(arr1)
	println(str2)
}

func TestTrimLeft(t *testing.T) {
	str := "corvus-rec-top-note-breakdown-realtime"
	// trimLeft有坑，只截取unicode， 特殊符号不截取
	left := strings.TrimLeft(str, "corvus-")
	println(left)
	left2 := strings.TrimPrefix(str, "corvus-")
	println(left2)
}

func TestString2Int(t *testing.T) {
	str1 := ""
	i1, err := strconv.Atoi(str1)
	if err != nil {
		println(err.Error())
		return
	}
	println(i1)
}

// 算法最长公共前缀
func TestLongestCommonPrefix(t *testing.T) {
	strs := []string{"flower", "flow", "flight"}
	result := findLongestCommonPrefix(strs)
	fmt.Println(result) // 输出 "hello"
}

func findLongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
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
