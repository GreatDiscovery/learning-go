package main

import (
	"errors"
	"fmt"
	pkg_errors "github.com/pkg/errors"
	"testing"
)

func TestErrorStack(t *testing.T) {
	err := causeStackError()
	if err != nil {
		fmt.Printf("%+v\n", err) // %+v 格式化会打印堆栈信息
	}

	fmt.Println("-------------------------------")

	err = causeNormalError()
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
}

func causeStackError() error {
	return pkg_errors.New("error with stack trace")
}

func causeNormalError() error {
	return errors.New("normal error")
}
