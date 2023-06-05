package main

import (
	"fmt"
	"testing"
)

// define interface and implementation

type Phone interface {
	call()
}

type Vivo struct{}
type Iphone struct{}

func (vivo Vivo) call() {
	fmt.Println("vivo call phone!")
}

func (apple Iphone) call() {
	fmt.Println("apple iphone call!")
}

func TestInterface(t *testing.T) {
	var phone Phone
	phone = new(Vivo)
	phone.call()

	phone = new(Iphone)
	phone.call()
}

func main() {

}
