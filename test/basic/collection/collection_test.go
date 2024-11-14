package main

import (
	"fmt"
	"github.com/go-playground/assert/v2"
	"sort"
	"testing"
)

type StringWalker func(s string) bool

func TestSort2(t *testing.T) {
	testCase := []string{"0.1234", "0.5789", "0.0000", "0.3232", "0.1233", "0.1212"}
	sort.Strings(testCase)
	fmt.Printf("sort list = %v\n", testCase)
	fmt.Printf("last element = %v\n", testCase[len(testCase)-1])
}

func TestSort(t *testing.T) {
	testCase := []string{"1", "2", "3", "2", "4", "4"}
	sort.Strings(testCase)
	fmt.Printf("sort list = %v\n", testCase)

	testCase2 := []float64{5, 7, 2, 1, 8, 6, 3, 4}
	testCase3 := sort.Float64Slice(testCase2)
	sort.Sort(testCase3)
	fmt.Printf("case3 = %v\n", testCase3)

	fmt.Printf("case2=%v\n", testCase2)
	sort.Sort(sort.Reverse(sort.Float64Slice(testCase2)))
	fmt.Printf("reverse case4=%v\n", testCase2)
}

func WalkList(strList []string, walker StringWalker) {
	for _, s := range strList {
		if walker(s) {
			println(s)
			return
		}
	}
}

func TestWalker(t *testing.T) {
	testCase := []string{"6", "2", "3", "2", "4", "4"}
	WalkList(testCase, func(s string) bool {
		if s == "3" {
			return true
		}
		return false
	})
}

func TestSet(t *testing.T) {
	testCase := []string{"1", "2", "3", "2", "4", "4"}
	set := make(map[string]bool)
	for _, s := range testCase {
		_, ok := set[s]
		if ok {
			continue
		}
		set[s] = true
	}
	var list []string
	for k, _ := range set {
		list = append(list, k)
	}
	fmt.Println(list)
}

func TestSliceDeleteItem(t *testing.T) {
	var arr = []int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); {
		if arr[i]%2 == 0 {
			// ...代表可变参数
			arr = append(arr[:i], arr[i+1:]...)
		} else {
			i++
		}
	}
	fmt.Printf("after del:%v", arr)
}

func TestSlice(t *testing.T) {
	var x = make([]int, 3, 5)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
	testCase := []string{"1", "2", "3", "2", "4", "4"}
	cutSlice := testCase[2]
	fmt.Printf("cutSlice=%v", cutSlice)
}

// 在切片前面添加新元素
func TestInsertIndex0(t *testing.T) {
	original := []int{2, 3, 4}
	newElement := 1

	updated := append([]int{newElement}, original...)
	assert.Equal(t, updated, []int{1, 2, 3, 4})
	fmt.Println(updated) // 输出: [1 2 3 4]
}

// len和cap的区别
func TestSliceCap(t *testing.T) {
	oldSlice := []int{0, 1, 2, 3, 4}
	assert.Equal(t, len(oldSlice), 5)
	assert.Equal(t, cap(oldSlice), 5)

	newSlice := oldSlice[1:3]
	assert.Equal(t, len(newSlice), 2)
	// 等于从索引 1 到底层数组末尾的元素数
	assert.Equal(t, cap(newSlice), 4)
	assert.Equal(t, newSlice, []int{1, 2})

	// 访问数组越界
	assert.PanicMatches(t, func() {
		fmt.Println("newSlice[3]=", newSlice[3])
	}, "runtime error: index out of range [3] with length 2")

	// newSlice不能随意更改老的slice
	newSlice = append(newSlice, 5)
	assert.Equal(t, len(oldSlice), 5)
	assert.Equal(t, oldSlice, []int{0, 1, 2, 5, 4})
	assert.Equal(t, oldSlice[3], 5)
	assert.Equal(t, newSlice, []int{1, 2, 5})
}
