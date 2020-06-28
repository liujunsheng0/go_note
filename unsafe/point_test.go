// 谨慎使用指针操作

package unsafe

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

// 通过指针操作修改栈上变量的值
func TestPointBaseTypeOperate(t *testing.T) {
	var a, b int
	// 获取b的地址
	ptr := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&a)) + unsafe.Sizeof(a)))
	// 修改b的值
	*ptr = 1

	assert.Equal(t, "*int", reflect.TypeOf(&a).String())
	assert.Equal(t, "unsafe.Pointer", reflect.TypeOf(unsafe.Pointer(&a)).String())
	assert.Equal(t, "uintptr", reflect.TypeOf(uintptr(unsafe.Pointer(&a))).String())

	assert.Equal(t, &b, ptr)
	assert.Equal(t, 0, a, "!=0")
	assert.Equal(t, 1, b)
	assert.Equal(t, 1, *ptr)
}

// 利用指针改变拥有公共组合的结构体
func TestPointStructOperate(t *testing.T) {
	// 如果把结构体中A 和变量顺序调换, 结果会不同
	a := A{}
	b := B{}
	c := C{}

	// 修改结构体中首个int变量的值
	// 谨慎使用此种方式
	setA((*A)(unsafe.Pointer(&a)), 1)
	setA((*A)(unsafe.Pointer(&b)), 1)
	setA((*A)(unsafe.Pointer(&c)), 1) // c中修改的是c.c的值, 因为修改的是c.c的地址

	assert.Equal(t, 1, a.a)
	assert.Equal(t, 1, b.a)
	assert.Equal(t, 1, c.c)

	// 修改a的值
	a.setA(2)
	b.setA(2)
	c.setA(2)

	assert.Equal(t, 2, a.a)
	assert.Equal(t, 2, b.a)
	assert.Equal(t, 2, c.a)
}

// 指针偏移
func TestPointOffset(t *testing.T) {
	b := B{}
	// 指针偏移, 需要转换为uintptr
	// 将指针偏移到b.b              起始位置                        b.b的位置
	ptr := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&b)) + unsafe.Offsetof(b.b)))
	// 修改b的值
	*ptr = 1
	assert.Equal(t, 1, b.b)
	assert.Equal(t, 1, *ptr)

	b.b = 22
	assert.Equal(t, 22, b.b)
	assert.Equal(t, 22, *ptr)

	// 指针偏移到b.b, 类似于C++的方式
	ptr = (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&b)) + unsafe.Sizeof(b.A)))
	*ptr = 333
	assert.Equal(t, 333, *ptr)
	assert.Equal(t, 333, b.b)

	a := A{}
	assert.Equal(t, 0, a.a)
	ptr = (*int)(unsafe.Pointer(&a))
	*ptr = 4444
	assert.Equal(t, 4444, *ptr)
	assert.Equal(t, 4444, a.a)
}
