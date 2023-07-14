package main

import (
	"fmt"
	"testing"
)

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
}
