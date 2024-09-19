package keyword

import (
	"fmt"
	"testing"
)

//fallthrough 的使用规则：
//fallthrough 只能用于 switch 语句中。
//它必须是某个 case 代码块的最后一条语句。
//使用 fallthrough 后，会无条件地执行下一个 case 代码块（即使下一个 case 的条件不匹配）

func TestFallThrough(t *testing.T) {
	num := 2
	switch num {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
		fallthrough
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	default:
		fmt.Println("Other")
	}
}
