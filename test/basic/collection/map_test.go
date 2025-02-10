package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type person struct {
	Name string
	Age  int
}

func (p person) string() string {
	return fmt.Sprintf("name=%s, age=%v", p.Name, p.Age)
}

func TestMapEqual(t *testing.T) {
	m1 := map[string]int{"a": 1, "b": 2}
	m2 := map[string]int{"a": 1, "b": 2}
	assert.Equal(t, m1, m2)
	m3 := map[string]int{"a": 1, "b": 3}
	assert.Equal(t, m1, m3)
}

func TestMapPointer(t *testing.T) {
	m := make(map[string]*person)
	if v, ok := m["k1"]; ok {
		fmt.Println("ok, v", v)
	} else {
		fmt.Println("not ok, v", v)
	}
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

	// 删除key
	delete(siteMap, "Google")
	fmt.Println("长度是", len(siteMap))
	name, ok = siteMap["Google"]
	assert.Equal(t, "", name)
}
