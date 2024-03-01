package main

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeFormat(t *testing.T) {
	// 获取当前时间
	currentTime := time.Now()
	// 设置时区
	location, _ := time.LoadLocation("Asia/Shanghai")
	// 将时间转换为指定时区
	currentTime = currentTime.In(location)
	// 格式化时间为指定的字符串格式
	formattedTime := currentTime.Format("2006-01-02T15:04:05-07:00")

	// 输出格式化后的时间字符串
	fmt.Println("格式化后的时间字符串:", formattedTime)
}
