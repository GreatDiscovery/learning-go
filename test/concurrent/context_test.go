package concurrent

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

// article :https://segmentfault.com/a/1190000040917752

type otherContext struct {
	context.Context
}

func TestTimeOut(t *testing.T) {
	root := context.Background()
	tm := time.Now().Add(3 * time.Second)
	ctxb, cancel := context.WithDeadline(root, tm)
	defer cancel()

	ch := make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		time.Sleep(2 * time.Second)
		ch <- struct{}{}
	}()
	go func(ctx context.Context) {
		defer wg.Done()
		select {
		case <-ctx.Done():
			fmt.Printf("get msg to cancel")
			return
		case <-ch:
			fmt.Println("work is done")
		}
	}(ctxb)
	wg.Wait()

}

func TestMultiValue(t *testing.T) {
	root := context.Background()
	c1 := context.WithValue(root, "key1", "value1")
	c2 := context.WithValue(c1, "key2", "value2")
	c3 := context.WithValue(c2, "key3", "value3")

	println(c3.Value("key1").(string))
	println(c3.Value("key2").(string))
	println(c3.Value("key3").(string))
}

func TestWithValue(t *testing.T) {
	type key struct{}
	root := context.Background()
	sub := context.WithValue(root, key{}, "key")
	s, _ := sub.Value(key{}).(string)
	fmt.Printf("s=%s\n", s)
}

func TestContext(t *testing.T) {
	ctxa, cancel := context.WithCancel(context.Background())
	go work(ctxa, "work1")

	tm := time.Now().Add(3 * time.Second)
	ctxb, _ := context.WithDeadline(ctxa, tm)
	go work(ctxb, "work2")

	oc := otherContext{ctxb}
	ctxc := context.WithValue(oc, "key", "andes,pass from main")
	go workWithValue(ctxc, "work3")

	time.Sleep(10 * time.Second)
	// 显示调用work1的cancel方法通知其退出
	cancel()
	fmt.Println("after cancel")

	// 等待work1打印退出信息
	time.Sleep(5 * time.Second)
	fmt.Println("main stop")

}

func work(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s get msg to cancel\n", name)
			return
		default:
			fmt.Printf("%s is running \n", name)
			time.Sleep(1 * time.Second)
		}
	}
}

func workWithValue(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s get msg to cancel\n", name)
			return
		default:
			value := ctx.Value("key").(string)
			fmt.Printf("%s is running value = %s \n", name, value)
			time.Sleep(1 * time.Second)
		}
	}
}
