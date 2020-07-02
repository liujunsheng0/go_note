package _return

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 返回值为局部变量指针: Go中返回局部变量指针是安全的, 编译器会根据需要将其分配在 GC Heap / stack 上
func TestReturnLocalVariablePoint(t *testing.T) {
	f := func() *int {
		a := 1
		return &a
	}
	p := f()
	assert.Equal(t, reflect.Ptr, reflect.TypeOf(p).Kind())
	assert.Equal(t, reflect.Int, reflect.TypeOf(*p).Kind())
	assert.Equal(t, 1, *p)

	*p = 2
	assert.Equal(t, 2, *p)
}

// 命名返回值
func TestReturnNamedVariable(t *testing.T) {
	// 函数返回值中的a和b在函数调用时就已经被赋予了一个初始零值
	// 即使函数使用了命名返回值, 依旧可以无视它而返回明确的值
	// 尽量使用命名返回值:会使代码更清晰、更简短, 同时更加容易读懂
	f := func(i bool) (a int, b string) {
		if i {
			a = 1
			b = "b"
		}
		return
	}
	a1, b1 := f(false)
	a2, b2 := f(true)
	assert.Equal(t, a1, 0)
	assert.Equal(t, b1, "")
	assert.Equal(t, a2, 1)
	assert.Equal(t, b2, "b")

}

// 返回值为函数
func TestReturnFunction(t *testing.T) {
	f := func(a, b int) func(int) int {
		// 闭包
		return func(c int) int {
			return a + b + c
		}
	}
	assert.Equal(t, 6, f(1, 2)(3))
}
