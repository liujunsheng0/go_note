package empty

import (
	"fmt"
	"testing"
)

// 空接口 == interface{}
type Any interface{}

// 实现类似于Python中的list
type List struct {
	element []Any
}

// List
// 接收者均为引用
func (l *List) Len() int {
	return len(l.element)
}

func (l *List) At(i int) Any {
	if i < 0 {
		i = l.Len() + i
	}
	return l.element[i]
}

func (l *List) Set(i int, e Any) {
	l.element[i] = e

}

func (l *List) Pop() Any {
	m := l.At(-1)
	l.element = l.element[:l.Len()-1]
	return m
}

func (l *List) Append(i Any) {
	l.element = append(l.element, i)
}

func (l *List) Insert(index int, i Any) {
	length := l.Len()
	if index >= (length - 1) {
		l.Append(i)
		return
	}
	for ; index < length; index++ {
		tmp := l.At(index)
		l.Set(index, i)
		i = tmp
	}
	l.Append(i)
}

func (l *List) Print(t *testing.T) {
	t.Log("index  value")
	for i, j := range l.element {
		t.Log(fmt.Sprintf("  %-2v     %v", i, j))
	}
}
