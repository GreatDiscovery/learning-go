package alogrithm

import (
	"github.com/joaojeronimo/go-crc16"
	"testing"
)

func TestCrc16(t *testing.T) {
	println(GetCRC16("_redis_rmt_check_152689"))
}

// 根据key算出redis的slot
func GetCRC16(str string) int {
	return int(crc16.Crc16([]byte(str))) % 16384
}
