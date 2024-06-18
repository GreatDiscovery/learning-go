package concurrent

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) {
	var a int64 = 0
	var b int64 = 0
	wg := sync.WaitGroup{}
	for i := 0; i < 500; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&a, 1)
		}()
	}
	wg.Wait()

	for i := 0; i < 500; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			b = b + 1
		}()
	}
	wg.Wait()
	fmt.Println("a=", a)
	fmt.Println("b=", b)
}
