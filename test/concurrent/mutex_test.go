package main

//https://laravelacademy.org/post/19928
import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

var counter int = 0

func add(a, b int, lock *sync.Mutex) {
	c := a + b
	lock.Lock()
	counter++
	fmt.Printf("%d: %d + %d = %d\n", counter, a, b, c)
	lock.Unlock()
}

func TestMutexLock(t *testing.T) {
	start := time.Now()
	//define a lock
	lock := &sync.Mutex{}
	for i := 0; i < 10; i++ {
		go add(1, i, lock)
	}

	for {
		lock.Lock()
		c := counter
		lock.Unlock()
		// 主线程暂时让渡cpu时间片给其他的goroutine不挂起，一段时间后会自动执行
		runtime.Gosched()
		if c >= 10 {
			break
		}
	}
	end := time.Now()
	consume := end.Sub(start).Seconds()
	fmt.Println("程序执行耗时(s)：", consume)
}
