package parameter_pass

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

// 值传递 引用传递
// slice, map, interface, channel 引用传递
func TestPassValueOrReference(t *testing.T) {
	// map - 引用传递
	{
		// a 和 b 指向的是同一个对象, 虽然指针值不同, 但是引用的对象是一样的
		a := map[string]int{"one": 1}
		b := a
		a["zero"] = 0
		assert.Equal(t, true, reflect.DeepEqual(a, b))
		assert.NotEqual(t, unsafe.Pointer(&a), unsafe.Pointer(&b))
		assert.Equal(t, &a, &b)
		func(m map[string]int) {
			m["two"] = 2
		}(a)
		assert.Equal(t, true, reflect.DeepEqual(a, b))
		assert.Equal(t, 2, a["two"])
		assert.Equal(t, &a, &b)
	}
	// slice - 引用传递
	{
		// a 和 b 指向的是同一个对象, 虽然指针不同, 但是引用的对象是一样的
		a := make([]int, 5)
		b := a
		assert.Equal(t, &a, &b)
		assert.Equal(t, &a[0], &b[0])
		assert.Equal(t, true, reflect.DeepEqual(a, b))
		// 注意append, 使用后地址就变了
		func(m []int) {
			size := cap(m)
			for i := 0; i < size; i++ {
				m[i] = i
			}
			assert.Equal(t, &m, &a)
			// 此后的修改不会影响外侧
			m = append(m, size)
			m[0] = 100
			assert.NotEqual(t, &m, &a)
		}(a)
		assert.Equal(t, true, reflect.DeepEqual(a, b))
	}
	// array - 值传递
	{
		a := [5]int{}
		b := a
		assert.Equal(t, false, &a[0] == &b[0])
		assert.Equal(t, true, reflect.DeepEqual(a, b))
		a[0] = 1
		assert.Equal(t, 1, a[0])
		assert.Equal(t, 0, b[0])

		func(m [5]int) {
			m[0] = 2
		}(a)
		assert.Equal(t, 1, a[0])
	}
}

// 变长参数
func TestPassVariableParameter(t *testing.T) {
	// b是变长参数, 类型是slice
	f := func(a ...int) {
		assert.LessOrEqual(t, 0, len(a))
		assert.Equal(t, reflect.Slice, reflect.TypeOf(a).Kind())
	}
	f()
	f(1, 2, 3)
}

// 函数作为参数
func TestPassFunction(t *testing.T) {
	add := func(a, b int) int {
		return a + b
	}

	sub := func(a, b int) int {
		return a - b
	}

	f := func(a, b int, callback func(int, int) int) int {
		return callback(a, b)
	}
	assert.Equal(t, 2, f(1, 1, add))
	assert.Equal(t, 0, f(1, 1, sub))
}
