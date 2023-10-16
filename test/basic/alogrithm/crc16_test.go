package alogrithm

import (
	"fmt"
	"github.com/joaojeronimo/go-crc16"
	"strconv"
	"testing"
	"time"
)

// 每个slot 0-16383分配一个key
func TestGetAllSlotKey(t *testing.T) {
	slotMap := make(map[int]string)
	prefix := "slotKey"
	now := strconv.Itoa(int(time.Now().Unix()))

	start := time.Now()
	const maxInt = 163850

	for i := 0; i < maxInt; i++ {
		if len(slotMap) >= 16384 {
			break
		}
		key := prefix + now + "-" + strconv.Itoa(i)
		slot := GetCRC16(key)
		if _, ok := slotMap[slot]; !ok {
			slotMap[slot] = key
		}
	}

	end := time.Now()
	fmt.Println("cost=", end.Sub(start))
	fmt.Println("slotMap=", slotMap)
	fmt.Println("slotMap[0]=", slotMap[0])
}

func TestCrc16(t *testing.T) {
	println(GetCRC16("slotKey1697441090-211"))
}

// 根据key算出redis的slot
func GetCRC16(str string) int {
	return int(crc16.Crc16([]byte(str))) % 16384
}
