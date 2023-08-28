package main

import (
	"fmt"
	"testing"
	"time"
)

func TestDate(t *testing.T) {
	now := time.Now()
	fmt.Printf("now=%v\n", now)
	yesterday := now.AddDate(0, 0, -1)
	fmt.Printf("yesterday=%v\n", yesterday)
	fmt.Printf("now.unix=%v, yesterday.unix=%v", now.Unix(), yesterday.Unix())
}
