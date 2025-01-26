package main

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       // 解析参数，默认是不会解析的
	fmt.Println(r.Form) // 这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") // 这个写入到 w 的是输出到客户端的
}

func TestWebServer(t *testing.T) {
	//访问http://localhost:9090
	http.HandleFunc("/", sayHelloName)       // 设置访问的路由
	err := http.ListenAndServe(":9090", nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func TestMockWebServer(t *testing.T) {
	// 创建一个模拟的 HTTP 服务器
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("pong"))
	}))
	defer server.Close() // 在测试完成后关闭服务器

	ping, err := Ping(server.URL)
	assert.NoError(t, err)
	assert.Equal(t, "pong", ping)
}

func Ping(url string) (string, error) {
	if url == "" {
		return "", errors.New("url is empty")
	}

	fullUrl := fmt.Sprintf("%s/ping", url)
	method := "GET"

	client := &http.Client{Timeout: time.Second * 1}
	req, err := http.NewRequest(method, fullUrl, nil)

	if err != nil {
		return "", err
	}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	if string(body) != "pong" {
		return "", errors.New(fmt.Sprintf("expect pong instead of %s", string(body)))
	}
	return string(body), nil
}
