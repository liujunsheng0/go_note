package nil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://studygolang.com/articles/10635  interface 与 nil 的比较
// https://studygolang.com/articles/18941  interface底层实现
// 在Go语言中, 一个interface{}类型的变量包含了2个指针, 一个指针指向值的类型, 另外一个指针指向实际的值
// 对于一个interface{}类型的nil变量来说, 它的两个指针都是0. 这是符合Go语言对nil的标准定义的. 在Go语言中，nil是零值.
// 将一个具体类型的值赋值给一个interface类型的变量的时候, 就同时把类型和值都赋值给了interface里的两个指针.
// 如果这个具体类型的值是nil的话, interface变量依然会存储对应的类型指针和值指针.
//
func TestNil(t *testing.T) {
	var pDog *Dog = nil
	var dog Dog

	var i1 Animal
	var i2 Animal
	var i3 Animal
	var i4 Animal

	i1 = nil // 类型和值均为0
	// 将*Dog类型的nil赋值给interface{},实际上interface里依然存了指向类型的指针, 所以拿这个interface变量去和nil常量进行比较的话就会返回false
	i2 = pDog // 类型不为0, 值为0
	i3 = &dog // 类型和值均不为0
	i4 = &dog // 类型和值均不为0

	t.Logf("i1 %v %+v\n", i1, asInterfaceStructure(i1))
	t.Logf("i2 %v %+v\n", i2, asInterfaceStructure(i2))
	t.Logf("i3 %v %+v\n", i3, asInterfaceStructure(i3))
	t.Logf("i4 %v %+v\n", i4, asInterfaceStructure(i4))

	assert.Equal(t, "", i2.Name())
	assert.Equal(t, "dog", i3.Name())
	assert.Equal(t, "dog", i4.Name())

	// 用以下方式判断接口类型是否为nil是错误的
	assert.Equal(t, true, i1 == nil)
	assert.Equal(t, false, i2 == nil)
	assert.Equal(t, false, i3 == nil)
	assert.Equal(t, false, i4 == nil)

	// reflect.ValueOf(i1).IsNil() 类型错误会抛出异常
	// 判断是否为空接口的正确方法如下
	assert.Equal(t, true, IsNil(nil))
	assert.Equal(t, true, IsNil(i1))
	assert.Equal(t, true, IsNil(i2))
	assert.Equal(t, false, IsNil(i3))
	assert.Equal(t, false, IsNil(i4))
	assert.Equal(t, false, IsNil(1))
	assert.Equal(t, false, IsNil("1"))
	assert.Equal(t, false, IsNil(1.1))

}
