package cronjob

import (
	"fmt"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	// 创建一个定时器，每隔1秒触发一次
	timer := time.NewTimer(1 * time.Second)

	// 创建一个无限循环，等待定时器触发事件
	go func() {
		count := 0
		for {
			select {
			case <-timer.C:
				// 定时器触发事件，执行相应的操作
				count = count + 1
				fmt.Println("定时器触发了！count=", count)
				// 重置定时器，使其再次触发
				timer.Reset(1 * time.Second)
			}
		}
	}()
	time.Sleep(6 * time.Second)
	timer.Stop()
}
