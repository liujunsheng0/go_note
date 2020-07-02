package map_

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	m = map[int]string{1: "one", 2: "two"}
)

// 在声明的时候不需要知道 map 的长度, map 是可以动态增长的
// 声明: var varName[key]value 引用类型
// 未初始化的 map 的值是 nil
// 通过 key 在 map 中寻找值是很快的, 比线性查找快得多, 但是仍然比从数组和切片的索引中直接读取要慢 100 倍
func TestInit(t *testing.T) {
	// 此时m1=nil, 不能进行操作
	var m1 map[int]string

	m2 := map[int]string{}
	m3 := make(map[int]string, 4)

	// 不要使用new初始化map, m4=nil
	m4 := *(new(map[int]string))

	// m1[0] = ""   // panic: assignment to entry in nil map
	assert.Nil(t, m1)
	// 添加成员
	m2[0] = "0"
	m3[0] = "0"
	// m4[0] = ""   // panic: assignment to entry in nil map
	assert.Nil(t, m4)
	t.Log(fmt.Sprintf("len(m1):%v m1:%v", len(m1), m1))
	t.Log(fmt.Sprintf("len(m2):%v m2:%v", len(m2), m2))
	t.Log(fmt.Sprintf("len(m3):%v m3:%v", len(m3), m3))
	t.Log(fmt.Sprintf("len(m3):%v m3:%v", len(m4), m4))
}

func TestGet(t *testing.T) {
	// 如果 map 中不存在 key, 返回对应类型的空值
	assert.Equal(t, "", m[0])
	// 存在found=true, 不存在为false
	v, found := m[0]
	assert.Equal(t, false, found)
	assert.Equal(t, "", v)

	assert.Equal(t, "one", m[1])
	v, found = m[1]
	assert.Equal(t, true, found)
	assert.Equal(t, "one", v)
}

func TestDelete(t *testing.T) {
	// 如果 key 不存在, 该操作不会产生错误
	delete(m, 0)
	m[0] = "zero"
	assert.Equal(t, "zero", m[0])
	delete(m, 0)
	assert.Equal(t, "", m[0])
}

func TestRange(t *testing.T) {
	for k := range m {
		t.Log(k)
	}

	for k, v := range m {
		t.Log(k, "->", v)
	}
}

func TestSortKey(t *testing.T) {
	m[0] = "zero"
	m[100] = "hundred"
	var keys []int
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		t.Log(fmt.Sprintf("%-3v -> %v", k, m[k]))
	}
}

func TestSwapKV(t *testing.T) {
	a := map[string]int{}
	for k, v := range m {
		a[v] = k
	}
	// a
	// one -> 1
	// two -> 2
	for k, v := range m {
		assert.Equal(t, k, a[v])
	}
}
