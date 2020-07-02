package size

import (
	"fmt"

	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

// 基本数据类型大小, 对应指针大小
func TestBaseTypeSize(t *testing.T) {
	var boolV bool

	var uint8V uint8
	var uint16V uint16
	var uint32V uint32
	var uint64V uint64

	var int8V int8
	var int16V int16
	var int32V int32
	var int64V int64
	var intV int

	var float32V float32
	var float64V float64
	structV := struct{ int }{}

	var interfaceV interface{}
	var strV string

	var mapV map[int]int
	var chanV chan int

	// 64位系统, 指针大小为8字节, 32位系统, 指针大小为4字节
	// sizeOf单位是byte
	// 地址 由低 -> 高
	t.Log("type  sizeOf     addr        sizeOf(point)")
	t.Log(fmt.Sprintf("bool        %-2d   %p      %d", unsafe.Sizeof(boolV), &boolV, unsafe.Sizeof(&boolV)))          // 1    0xc000020358      8
	t.Log(fmt.Sprintf("uint8       %-2d   %p      %d", unsafe.Sizeof(uint8V), &uint8V, unsafe.Sizeof(&uint8V)))       // 1    0xc000020359      8
	t.Log(fmt.Sprintf("uint16      %-2d   %p      %d", unsafe.Sizeof(uint16V), &uint16V, unsafe.Sizeof(&uint16V)))    // 2    0xc00002035a      8
	t.Log(fmt.Sprintf("uint32      %-2d   %p      %d", unsafe.Sizeof(uint32V), &uint32V, unsafe.Sizeof(&uint32V)))    // 4    0xc00002035c      8
	t.Log(fmt.Sprintf("uint64      %-2d   %p      %d", unsafe.Sizeof(uint64V), &uint64V, unsafe.Sizeof(&uint64V)))    // 8    0xc000020360      8
	t.Log(fmt.Sprintf("int8        %-2d   %p      %d", unsafe.Sizeof(int8V), &int8V, unsafe.Sizeof(&int8V)))          // 1    0xc000020368      8
	t.Log(fmt.Sprintf("int16       %-2d   %p      %d", unsafe.Sizeof(int16V), &int16V, unsafe.Sizeof(&int16V)))       // 2    0xc00002036a      8
	t.Log(fmt.Sprintf("int32       %-2d   %p      %d", unsafe.Sizeof(int32V), &int32V, unsafe.Sizeof(&int32V)))       // 4    0xc00002036c      8
	t.Log(fmt.Sprintf("int64       %-2d   %p      %d", unsafe.Sizeof(int64V), &int64V, unsafe.Sizeof(&int64V)))       // 8    0xc000020370      8
	t.Log(fmt.Sprintf("int         %-2d   %p      %d", unsafe.Sizeof(intV), &intV, unsafe.Sizeof(&intV)))             // 8    0xc000020378      8
	t.Log(fmt.Sprintf("float32     %-2d   %p      %d", unsafe.Sizeof(float32V), &float32V, unsafe.Sizeof(&float32V))) // 4    0xc000020380      8
	t.Log(fmt.Sprintf("float64     %-2d   %p      %d", unsafe.Sizeof(float64V), &float64V, unsafe.Sizeof(&float64V))) // 8    0xc000020388      8
	t.Log(fmt.Sprintf("struct      %-2d   %p      %d", unsafe.Sizeof(structV), &structV, unsafe.Sizeof(&structV)))    // 8    0xc000020390      8

	t.Log(fmt.Sprintf("map         %-2d   %p      %d", unsafe.Sizeof(mapV), &mapV, unsafe.Sizeof(&mapV)))    // 8    0xc00000e038      8
	t.Log(fmt.Sprintf("chan        %-2d   %p      %d", unsafe.Sizeof(chanV), &chanV, unsafe.Sizeof(&chanV))) // 8    0xc00000e040      8

	t.Log(fmt.Sprintf("interface{} %-2d   %p      %d", unsafe.Sizeof(interfaceV), &interfaceV, unsafe.Sizeof(&interfaceV))) // 16   0xc00005c4d0      8
	t.Log(fmt.Sprintf("string      %-2d   %p      %d", unsafe.Sizeof(strV), &strV, unsafe.Sizeof(&strV)))                   // 16   0xc00005c4e0      8
}

// 结构体大小
// 内存对齐
// func Alignof(x ArbitraryType) uintptr
//    接受任何类型的表达式x并返回假设变量x所需的对齐方式,
//    如果变量s是结构类型, 而f是该结构中的字段, 则Alignof(s.f)将返回结构中该类型字段的所需对齐方式
// 内存布局中一般是8字节对齐(不足八字节, 会补齐), 所以排列结构体时候要注意排列顺序
// 具体的对齐方式由结构体中的变量类型而定, 一般为:min(SizeOf(fields), 8)
func TestStructSize(t *testing.T) {
	s1 := struct {
		A bool
	}{}
	s2 := struct {
		A int16
	}{}
	s3 := struct {
		A int16
		B bool
	}{}
	s4 := struct {
		A int32
	}{}
	s5 := struct {
		A int32
		B bool
	}{}
	s6 := struct {
		A int32
		B int16
	}{}
	s7 := struct {
		A int32
		B int16
		C bool
	}{}
	// 注意s8a s8b s8c
	s8a := struct {
		A bool
		b bool
		c bool
		d bool
		e bool
		f bool
		g bool
		h bool
	}{}
	s8b := struct {
		A int32
		B int16
		C int16
	}{}
	s8c := struct {
		A int64
	}{}
	s40 := struct {
		A bool
		B float64
		C bool
		D float64
		E bool
	}{}
	s24 := struct {
		A bool
		B bool
		C bool
		D float64
		E float64
	}{}
	at := assert.New(t)
	at.Equal(1, int(unsafe.Sizeof(s1)))
	at.Equal(1, int(unsafe.Alignof(s1)))
	at.Equal(2, int(unsafe.Sizeof(s2)))
	at.Equal(2, int(unsafe.Alignof(s2)))
	at.Equal(4, int(unsafe.Sizeof(s3)))
	at.Equal(2, int(unsafe.Alignof(s3)))
	at.Equal(4, int(unsafe.Sizeof(s4)))
	at.Equal(4, int(unsafe.Alignof(s4)))
	at.Equal(8, int(unsafe.Sizeof(s5)))
	at.Equal(4, int(unsafe.Alignof(s5)))
	at.Equal(8, int(unsafe.Sizeof(s6)))
	at.Equal(4, int(unsafe.Alignof(s6)))
	at.Equal(8, int(unsafe.Sizeof(s7)))
	at.Equal(4, int(unsafe.Alignof(s7)))
	at.Equal(8, int(unsafe.Sizeof(s8a)))
	at.Equal(1, int(unsafe.Alignof(s8a)))
	at.Equal(8, int(unsafe.Sizeof(s8b)))
	at.Equal(4, int(unsafe.Alignof(s8b)))
	at.Equal(8, int(unsafe.Sizeof(s8c)))
	at.Equal(8, int(unsafe.Alignof(s8c)))
	at.Equal(24, int(unsafe.Sizeof(s24)))
	at.Equal(8, int(unsafe.Alignof(s24)))
	at.Equal(40, int(unsafe.Sizeof(s40)))
	at.Equal(8, int(unsafe.Alignof(s40)))
}

func TestArraySize(t *testing.T) {
	var slice1 []int
	slice2 := []int{1, 2, 3}

	assert.Equal(t, 24, int(unsafe.Sizeof(slice1)))
	assert.Equal(t, 24, int(unsafe.Sizeof(slice2)))

	slice1 = append(slice1, 4)
	assert.Equal(t, 24, int(unsafe.Sizeof(slice1)))

	map1 := map[int]int{1: 1, 2: 2}
	map2 := map[int]int{}
	assert.Equal(t, 8, int(unsafe.Sizeof(map1)))
	assert.Equal(t, 8, int(unsafe.Sizeof(map2)))

	map1[3] = 3
	assert.Equal(t, 8, int(unsafe.Sizeof(map1)))
}
