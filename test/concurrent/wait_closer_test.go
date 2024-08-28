package concurrent

import (
	"fmt"
	"io"
	"sync"
	"testing"
)

// WaitCloser 结构体
type WaitCloser struct {
	wg sync.WaitGroup
	c  io.Closer
}

// Add 方法，用于在开始并发操作时调用
func (wc *WaitCloser) Add(delta int) {
	wc.wg.Add(delta)
}

// Done 方法，用于在并发操作完成时调用
func (wc *WaitCloser) Done() {
	wc.wg.Done()
}

// Close 方法，确保所有并发操作完成后关闭资源
func (wc *WaitCloser) Close() error {
	wc.wg.Wait()
	return wc.c.Close()
}

// 假设一个示例资源
type exampleResource struct{}

// 实现 io.Closer 接口的 Close 方法
func (er *exampleResource) Close() error {
	fmt.Println("Resource closed")
	return nil
}

func TestWaitCloser(t *testing.T) {
	resource := &exampleResource{}
	wc := WaitCloser{c: resource}

	// 开始一个并发操作
	wc.Add(1)
	go func() {
		defer wc.Done()
		fmt.Println("Concurrent operation 1")
	}()

	// 再开始另一个并发操作
	wc.Add(1)
	go func() {
		defer wc.Done()
		fmt.Println("Concurrent operation 2")
	}()

	// 关闭资源（等待所有操作完成后关闭）
	wc.Close()
}
