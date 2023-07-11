package main

import (
	"testing"
)

func TestIf(t *testing.T) {
	var str string
	str = ""
	if str != "" {
		print(true)
	}
}
