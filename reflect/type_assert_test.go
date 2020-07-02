// 类型断言和反射

package reflect

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// interface 类型转换
func TestTypeAssertion(t *testing.T) {
	defer Recover(t)
	var a interface{} = "s"
	// v1: 转换后的值
	// ok: 是否转换成功, bool类型
	v1, ok1 := a.(string)
	v2, ok2 := a.(int)
	assert.Equal(t, true, ok1)
	assert.Equal(t, "s", v1)
	assert.Equal(t, false, ok2)
	assert.Equal(t, 0, v2)

	// 直接转换, 如果出错抛出异常, 不建议使用
	assert.Equal(t, "s", a.(string))
	// 类型错误, 转换时直接抛出异常, interface conversion: interface {} is string, not int
	_ = a.(int)
}

// interface switch 类型转换
func TestTypeAssertionSwitch(t *testing.T) {
	help := func(i interface{}) {
		// v = 对应转换后的值, 传递给switch是对应的类型
		switch v := i.(type) {
		default:
			fmt.Printf("unexpected type: %s\n", reflect.TypeOf(i))
		case bool:
			fmt.Printf("bool  %t\n", v)
		case int:
			fmt.Printf("int   %d\n", v)
		case *bool:
			fmt.Printf("*bool %v %p\n", *v, v)
		case *int:
			fmt.Printf("*int  %v %p\n", *v, v)
		}
	}
	a := 1
	b := false
	c := ""
	help(a)
	help(&a)
	help(b)
	help(&b)
	help(c)
	help(&c)
}
