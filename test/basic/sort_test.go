package main

import (
	"fmt"
	"sort"
	"strings"
	"testing"
)

func TestSortByFunc(t *testing.T) {
	s1 := []string{"xx-1-0", "xx-2-1", "xx-51-0", "xx-51-1", "xx-3-1", "xx-3-0"}
	sort.Slice(s1, func(i, j int) bool {
		split1 := strings.Split(s1[i], "-")
		split2 := strings.Split(s1[j], "-")
		role1 := split1[len(split1)-1]
		role2 := split2[len(split2)-1]
		version1 := split1[len(split1)-2]
		version2 := split2[len(split2)-2]
		if role1 > role2 {
			return true
		}
		if version1 < version2 {
			return true
		}
		return false
	})
	fmt.Println(s1)
}
