package helm

import (
	"fmt"
	"os"
	"testing"
	"text/template"
)

// 语法：https://www.topgoer.com/%E5%B8%B8%E7%94%A8%E6%A0%87%E5%87%86%E5%BA%93/template.html

type Address struct {
	City  string
	State string
}

type User struct {
	Name    string
	Age     int
	Address Address
	Hobbies []string
	Scores  map[string]int
}

type ConfigUser struct {
	Users []User
}

func TestRenderComplexStruct(t *testing.T) {
	// 准备数据
	config := ConfigUser{
		Users: []User{
			{
				Name: "Alice",
				Age:  25,
				Address: Address{
					City:  "New York",
					State: "NY",
				},
				Hobbies: []string{"Reading", "Swimming"},
				Scores: map[string]int{
					"Math":    90,
					"Science": 85,
				},
			},
			{
				Name: "Bob",
				Age:  30,
				Address: Address{
					City:  "San Francisco",
					State: "CA",
				},
				Hobbies: []string{"Coding", "Gaming"},
				Scores: map[string]int{
					"Math":    95,
					"Science": 88,
				},
			},
		},
	}

	// 解析模板文件
	tmpl, err := template.ParseFiles("complex_template.tmpl")
	if err != nil {
		fmt.Println("Failed to parse template:", err)
		return
	}

	// 渲染模板
	if err := tmpl.Execute(os.Stdout, config); err != nil {
		fmt.Println("Failed to execute template:", err)
		return
	}
}
