package slice

import (
	"fmt"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func TestArrayInit(t *testing.T) {
	var a1 [2]int        // [2]int
	a2 := [...]int{0, 1} // [2]int  自动计算长度
	a3 := [...]int{1: 1} // [2]int  自动计算长度, 索引1处置为1
	t.Log("a1:", spew.Sdump(a1))
	t.Log("a2:", spew.Sdump(a2))
	t.Log("a3:", spew.Sdump(a3))
}

func TestSliceInit(t *testing.T) {
	var s1 []int             // []int
	s2 := []int{0, 1}        // []int  自动计算长度和容量
	s3 := []int{1: 1}        // []int  自动计算长度和容量, 索引1处置为1, 其余为默认值
	s4 := make([]int, 2)     // []int  长度为2, 容量为2
	s5 := make([]int, 0, 10) // []int  长度为0, 容量为10, 容量为可选参数
	s4 = append(s4, 1)
	s5 = append(s5, 1)
	t.Log("s1:", spew.Sdump(s1))
	t.Log("s2:", spew.Sdump(s2))
	t.Log("s3:", spew.Sdump(s3))
	t.Log("s4:", spew.Sdump(s4))
	t.Log("s4:", spew.Sdump(s5))
}

// 如果多个切片表示同一个数组的片段, 他们是共享内存的, 即改变共享内存中的数据, 会影响所有引用了此处的切片和相应的数组
func TestArrayAndSlice(t *testing.T) {
	// 数组容量声明后, 就不可变了
	s := [5]int{}
	// s的切片是对s的引用
	s1 := s[:3]
	s2 := s[1:5]
	s[1] = 1

	ar := assert.New(t)
	// addr
	ar.Equal(&s[1], &s1[1])
	ar.Equal(&s[1], &s2[0])
	// value
	ar.Equal(1, s[1])
	ar.Equal(1, s1[1])
	ar.Equal(1, s2[0])

	// 切片和array 容量
	ar.Equal(cap(s1), 5)
	ar.Equal(cap(s2), 4)

	// 容量足够, 会影响array
	// 重新申请内存, 则不会影响
	s1 = append(s1, 3)
	ar.Equal(s1[3], 3)
	ar.Equal(s[3], 3)
	s2 = append(s2, 5)
	ar.Equal(s2[4], 5)
	ar.Equal(s[4], 0)
}

func TestAppend(t *testing.T) {
	s := make([]int, 0, 2)
	t.Log(fmt.Sprintf("addr:%p len:%v cap:%v s:%v", s, len(s), cap(s), s))

	// 容量足够, 不会重新分配内存
	s = append(s, 0, 1)
	t.Log(fmt.Sprintf("addr:%p len:%v cap:%v s:%v", s, len(s), cap(s), s))

	// 容量不足, 重新分配内存
	s = append(s, 2, 3)
	t.Log(fmt.Sprintf("addr:%p len:%v cap:%v s:%v", s, len(s), cap(s), s))

	// 扩容规则  https://www.bilibili.com/video/BV1hv411x7we?p=2
	// 还需要匹配内存
	// if oldCap * 2 < 所需cap
	//    newCap = 所需cap
	// else:
	// 	  if 元素个数 < 1024:  newCap = oldCap * 2
	//    else:  newCap = oldCap * 1.25
	// 扩容后的容量还需要考虑扩容后内存管理部分
	// 内存规格：8, 16, 32, 48, 64, 80.... byte
	s = make([]int, 0, 1)
	s = append(s, 1, 2)
	ar := assert.New(t)
	ar.Equal(cap(s), 2)
	s = append(s, 3, 4, 5)
	// 5 * sizeof(int) = 40, 需要分配到48byte内存管理部分, 所以newCap = 6
	ar.Equal(cap(s), 6)
}
