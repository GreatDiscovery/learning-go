package concurrent

// how to use ant framework to manage and recycle goroutine

import (
	"context"
	"fmt"
	"github.com/panjf2000/ants/v2"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

var sum1 int32

func Test5(t *testing.T) {

	num := 1

	switch num {
	case 1:
		fmt.Println("This is the first number.")
	case 2:
		fmt.Println("This is the second number.")
	case 3:
		fmt.Println("This is the third number.")
	default:
		fmt.Println("This is not the first, second, or third number.")
	}

}

func myFunc(i interface{}) {
	n := i.(int32)
	atomic.AddInt32(&sum1, n)
	fmt.Printf("run with %d\n", n)
}

func demoFunc() {
	time.Sleep(10 * time.Millisecond)
	fmt.Println("Hello World!")
}

func TestAnts(t *testing.T) {
	defer ants.Release()

	runTimes := 1000

	// Use the common pool.
	var wg sync.WaitGroup
	syncCalculateSum := func() {
		demoFunc()
		wg.Done()
	}
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		_ = ants.Submit(syncCalculateSum)
	}
	wg.Wait()

	fmt.Printf("running goroutines: %d\n", ants.Running())
	fmt.Printf("finish all tasks.\n")

	// Use the pool with a function,
	// set 10 to the capacity of goroutine pool and 1 second for expired duration.
	p, _ := ants.NewPoolWithFunc(10, func(i interface{}) {
		myFunc(i)
		wg.Done()
	})
	defer p.Release()
	// Submit tasks one by one.
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		_ = p.Invoke(int32(i))
	}
	wg.Wait()
	fmt.Printf("running goroutines: %d\n", p.Running())
	fmt.Printf("finish all tasks, result is %d\n", sum1)
}

func TestAntsWithError(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	wg := sync.WaitGroup{}

	pool, err := ants.NewPoolWithFunc(10, func(i interface{}) {
		defer wg.Done()
		n := i.(int32)
		select {
		case <-ctx.Done():
			fmt.Println(fmt.Sprintf("context done, %d", n))
		default:
			if n == 10 {
				fmt.Println("context error,", n)
				cancel()
				return
			}
			fmt.Println(fmt.Sprintf("work doing, %d", n))
			time.Sleep(100 * time.Millisecond)
		}
	})
	if err != nil {
		panic(err)
	}
	defer pool.Release()

	for i := 0; i < 20; i++ {
		if ctx.Err() != nil {
			return
		}
		wg.Add(1)
		err := pool.Invoke(int32(i))
		time.Sleep(time.Millisecond * 100)
		if err != nil {
			panic(err)
		}
	}
	wg.Wait()
}
