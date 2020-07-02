package spew

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

// go-spew可以帮打印数据的结构, 可以看到一个变量的数据结构信息

func TestSpew(t *testing.T) {
	str := "1234"
	slice := []string{"1", "2", "3"}
	array := [4]string{}
	m := map[int]string{1: "1", 0: "0"}
	s := struct {
		A int
		B string
	}{}
	ch := make(chan string, 0)
	var i interface{} = str

	spew.Dump(str, slice, array, m, s, ch, i)
}
