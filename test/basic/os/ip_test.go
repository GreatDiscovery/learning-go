package main

import (
	"fmt"
	"net"
	"testing"
)

func TestParseIp(t *testing.T) {
	ip := net.ParseIP("1.2.3.4")
	fmt.Printf("ip=%v", ip)
}
