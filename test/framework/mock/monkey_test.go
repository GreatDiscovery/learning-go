package mock

import (
	"fmt"
	"github.com/agiledragon/gomonkey/v2"
	"testing"
)

func DoSomething() string {
	return "Real Result"
}

// go tool evn 加上这个参数才会生效-gcflags=all=-l
// fixme 程序很容易崩溃，gomonkey有bug
func TestCompute(t *testing.T) {
	// 使用 gomonkey 替换 DoSomething 函数的实现
	patch := gomonkey.ApplyFunc(DoSomething, func() string {
		return "Mocked Result"
	})
	defer patch.Reset() // 确保测试结束后恢复原始函数

	// 验证函数返回的是我们模拟的结果
	result := DoSomething()
	fmt.Println(result)
}
