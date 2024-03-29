package concurrent

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestSelect(t *testing.T) {
	ch := make(chan int)
	go func() {
		time.Sleep(5 * time.Second)
		ch <- 5
	}()

	for {
		select {
		case num := <-ch:
			fmt.Println("num=", num)
			return
		default:
			time.Sleep(1 * time.Second)
			fmt.Println("if no default, if will block")
		}
	}
}

// 模拟10个协程，同时处理任务，最后收集返回结果
func TestMultiThread(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		go func() {
			num := rand.Intn(10)
			ch <- num

		}()
	}

	// 由于没有close函数，chan会一直阻塞在这里
	for i := range ch {
		fmt.Println(fmt.Sprintf("value=%v", i))
	}

	fmt.Println(time.Now().Unix() - start.Unix())
}
