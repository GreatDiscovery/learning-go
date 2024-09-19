package keyword

import (
	"testing"
)

//defer is a keyword in golang
//https://tiancaiamao.gitbooks.io/go-internals/content/zh/03.4.html

func f1() (result int) {
	defer func() {
		result++
	}()
	return 0
}

func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

func TestF(t *testing.T) {
	println(f1())
	println(f2())
	println(f3())
}
