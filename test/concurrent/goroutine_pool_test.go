package main

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

// 该方法只是一个简单的协程池，后续项目里可以接住一些框架类实现
// CoroutinePool是一个协程池类型，它包含协程池的初始化和销毁函数。
type CoroutinePool struct {
	Name string
	// 协程池大小
	MaxSize int
	// 协程池任务队列
	Queue chan func()
	// 协程池线程数
	NumThreads int
	// 协程池管理器
	Manager *sync.Pool
}

// Start starts the coroutine pool.
func (p *CoroutinePool) Start() {
	p.Manager = &sync.Pool{
		New: func() interface{} {
			return &sync.WaitGroup{}
		},
	}
	for i := 0; i < p.NumThreads; i++ {
		wg := p.Manager.Get().(*sync.WaitGroup)
		wg.Add(1)
		go func() {
			defer wg.Done()
			for task := range p.Queue {
				if task != nil {
					p.Manager.Put(task)
				}
			}
		}()
	}
}

// Stop stops the coroutine pool.
func (p *CoroutinePool) Stop() {
	p.Manager.Put(nil)
	for i := 0; i < p.NumThreads; i++ {
		wg := p.Manager.Get().(*sync.WaitGroup)
		wg.Wait()
	}
}

// Submit submits a task to the coroutine pool.
func (p *CoroutinePool) Submit(task func()) {
	if p.Queue == nil {
		return
	}
	p.Queue <- task
}

// NewCoroutinePool returns a new coroutine pool.
func NewCoroutinePool(name string, maxSize int, numThreads int) *CoroutinePool {
	return &CoroutinePool{
		Name:       name,
		MaxSize:    maxSize,
		NumThreads: numThreads,
		Queue:      make(chan func(), maxSize),
	}
}

// fixme it doesn't work
func TestGoroutinePool(t *testing.T) {
	task := func() {
		fmt.Println("Task executed")
		fmt.Println("Task executed")
		return
	}
	p := NewCoroutinePool("CoroutinePool", 10, 5)
	p.Submit(task)
	p.Start()
	runtime.Gosched()
	time.Sleep(5 * time.Second)
	p.Stop()
}
