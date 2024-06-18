package main

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
)

// 测试并发调用rand冲突的概率
func TestRand(t *testing.T) {
	counter := make(map[int]int)

	wg := sync.WaitGroup{}
	var mu sync.Mutex

	for i := 0; i < 30; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 100000; i++ {
				randInt := rand.Int()
				mu.Lock()
				counter[randInt] = counter[randInt] + 1
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	for k, v := range counter {
		if v > 1 {
			fmt.Println(fmt.Sprintf("k=%d, v=%d", k, v))
		}
	}
	fmt.Println("end")

}
