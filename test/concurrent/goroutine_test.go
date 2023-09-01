package concurrent

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

import "sync"

// 如果存在多个协程，其中一个协程出现了错误，我们应该如何处理
func TestGoRoutineWithError(t *testing.T) {
	var wg sync.WaitGroup
	errors := make(chan error)

	for i := 1; i <= 5; i++ {
		wg.Add(1) // 增加等待组计数器
		go func(id int) {
			defer wg.Done() // 在协程结束时减少等待组计数器
			err := doSomething(id)
			if err != nil {
				errors <- err // 将错误发送到通道
			}
		}(i)
	}

	go func() {
		wg.Wait()     // 等待所有协程完成
		close(errors) // 关闭错误通道，表示没有更多的错误了
	}()

	for err := range errors {
		// 处理错误，只处理一次即可
		fmt.Println("Error:", err)
		break // 退出循环，不再接收更多的错误信息
	}
}

func doSomething(id int) error {
	if rand.Intn(2) == 0 {
		fmt.Println("normal")
		return nil // 没有错误发生，返回 nil
	} else {
		return fmt.Errorf("something went wrong") // 发生错误，返回错误
	}
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func TestGoRoutine(t *testing.T) {
	go say("hello")
	say("world")
}

// how to use channel for passing particular data type.

func sum(s []int, c chan int) {
	sum := 0
	for _, value := range s {
		sum += value
	}
	c <- sum
}

func TestChannel(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}

func TestChannelWithBuffer(t *testing.T) {
	c := make(chan int, 2)
	c <- 1
	c <- 2
	fmt.Println(<-c)
	fmt.Println(<-c)
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func TestChannelWithClose(t *testing.T) {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	// range 函数遍历每个从通道接收到的数据，因为 c 在发送完 10 个
	// 数据之后就关闭了通道，所以这里我们 range 函数在接收到 10 个数据
	// 之后就结束了。如果上面的 c 通道不关闭，那么 range 函数就不
	// 会结束，从而在接收第 11 个数据的时候就阻塞了。
	for i := range c {
		fmt.Println(i)
	}
}

// channel可以同步，比如阻塞等待main进程
func TestChannelSync(t *testing.T) {
	c := make(chan int)
	go sum2(c)
	// 读取channel，会一直阻塞
	<-c
}

func sum2(i chan int) {
	sum := 0
	for i := 0; i < 10000; i++ {
		sum += i
	}
	println(sum)
	i <- sum
}
