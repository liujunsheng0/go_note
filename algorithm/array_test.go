package algorithm

import (
	"fmt"
	"testing"
)

// 递归反转数组
func ReverseArray(arr []int) []int {
	if len(arr) == 0 {
		return nil
	}
	return append(ReverseArray(arr[1:]), arr[0])
}

func Test_ReverseArray(t *testing.T) {
	a := "abcdefg我"
	for _, v := range a {
		fmt.Println(string(v))
	}
}
