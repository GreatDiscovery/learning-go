package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"testing"
	"time"
)

func shuffle(arr []int) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
}

func TestShuffle(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println("Original array:", arr)

	shuffle(arr)
	fmt.Println("Shuffled array:", arr)
}

func TestSortByFunc(t *testing.T) {
	str1 := "k8redis-jiayun-simba-sh4-1-0-3 k8redis-jiayun-simba-sh4-1-0-2 k8redis-jiayun-simba-sh4-1-1-3 k8redis-jiayun-simba-sh4-1-2-3 k8redis-jiayun-simba-sh4-1-1-2 k8redis-jiayun-simba-sh4-1-2-2"
	s1 := strings.Split(str1, " ")
	fmt.Println("before sort =", s1)
	sortSlice(s1)
	fmt.Println("after sort =", s1)
}

func sortSlice(s1 []string) []string {
	sort.Slice(s1, func(i, j int) bool {
		split1 := strings.Split(s1[i], "-")
		split2 := strings.Split(s1[j], "-")
		role1 := split1[len(split1)-1]
		role2 := split2[len(split2)-1]
		version1 := split1[len(split1)-2]
		version2 := split2[len(split2)-2]
		if role1 > role2 {
			return true
		} else if role1 == role2 {
			if version1 < version2 {
				return true
			} else {
				return false
			}
		} else {
			return false
		}
	})
	return s1
}
