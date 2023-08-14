package main

import (
	"fmt"
	"sort"
	"testing"
)

type StringWalker func(s string) bool

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
}

func TestMap(t *testing.T) {
	var siteMap map[string]string /*创建集合 */
	siteMap = make(map[string]string)

	/* map 插入 key - value 对,各个国家对应的首都 */
	siteMap["Google"] = "谷歌"
	siteMap["Runoob"] = "菜鸟教程"
	siteMap["Baidu"] = "百度"
	siteMap["Wiki"] = "维基百科"

	/*使用键输出地图值 */
	for site := range siteMap {
		fmt.Println(site, "首都是", siteMap[site])
	}

	/*查看元素在集合中是否存在 */
	name, ok := siteMap["Facebook"] /*如果确定是真实的,则存在,否则不存在 */
	/*fmt.Println(capital) */
	/*fmt.Println(ok) */
	if ok {
		fmt.Println("Facebook 的 站点是", name)
	} else {
		fmt.Println("Facebook 站点不存在")
	}

	fmt.Println("长度是", len(siteMap))
}
