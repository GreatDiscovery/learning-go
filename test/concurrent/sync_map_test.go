package concurrent

import (
	"fmt"
	"sync"
	"testing"
)

func TestSyncMap(t *testing.T) {
	var m sync.Map
	var sg sync.WaitGroup
	m.Store("key", 1)
	for i := 0; i < 10; i++ {
		sg.Add(1)
		go func(n int) {
			defer sg.Done()
			value, ok := m.Load("key")
			fmt.Printf("%v get value = %v\n", n, value)
			if ok {
				m.Store("key", value.(int)+1)
			} else {
				fmt.Println("未获取到key")
			}
		}(i)
	}
	sg.Wait()
	value, _ := m.Load("key")
	fmt.Println(value)
}
