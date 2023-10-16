package main

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestPrintErrorStack(t *testing.T) {
	err := errors.New("hello world")
	// 这里err没有携带错误堆栈
	log.Error("print stack,", err)
	fmt.Printf("err=%v", err)
	panic(err)
}
