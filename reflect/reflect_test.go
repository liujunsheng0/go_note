package reflect

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	// 变量命名为类型首字母
	n  interface{} = nil
	s  interface{} = ""
	i  interface{} = 1
	b  interface{} = true
	f  interface{} = 1.1
	st interface{} = struct{}{}
	sl interface{} = []int{0}
	m  interface{} = map[int]int{}
)

// func TypeOf(i interface{}) Type 获取i的类型相关信息, 如果i为nil, 返回nil
func TestTypeOf(t *testing.T) {
	//assert.Equal(t, reflect.Interface, reflect.TypeOf(n).Kind())
	assert.Equal(t, reflect.String, reflect.TypeOf(s).Kind())
	assert.Equal(t, reflect.Int, reflect.TypeOf(i).Kind())
	assert.Equal(t, reflect.Bool, reflect.TypeOf(b).Kind())
	assert.Equal(t, reflect.Float64, reflect.TypeOf(f).Kind())
	assert.Equal(t, reflect.Struct, reflect.TypeOf(st).Kind())
	assert.Equal(t, reflect.Slice, reflect.TypeOf(sl).Kind())
	assert.Equal(t, reflect.Map, reflect.TypeOf(m).Kind())

}

func TestValueConversion(t *testing.T) {
	defer Recover(nil)
	v := reflect.ValueOf(f)

	v1, ok1 := v.Interface().(float64)
	v2, ok2 := v.Interface().(int)
	v3 := v.Float() // 类型不符, 将会panic
	assert.Equal(t, ok1, true)
	assert.Equal(t, ok2, false)
	assert.Equal(t, 1.1, v1)
	assert.Equal(t, 0, v2)
	assert.Equal(t, 1.1, v3)

	// panic call of reflect.Value.Int on float64 Value
	//v.Int()
}

// 结构体中包含的属性和方法
func TestStructContain(t *testing.T) {
	f := func(i interface{}) {
		v := reflect.ValueOf(i)

		// 仅适用于结构体实
		if v.Kind() == reflect.Struct {
			// 获取结构体中的属性
			for i := 0; i < v.NumField(); i++ { // NumField returns the number of fields in the struct v.
				tf := v.Type().Field(i) // StructField
				vf := v.Field(i)        // Value
				fmt.Printf("  %-7v : %-4v (%-6v) index:%v offset:%v\n", tf.Name, vf, tf.Type, tf.Index, tf.Offset)
			}
		}

		// 获取结构体方法
		// 注意: 接收者为实例和指针的区别
		// typeOf(i) 结构体实例: 仅能获取接收者为实例的方法
		// typeOf(i) 结构体指针: 能获取结构体中的所有方法(包括实例的方法)
		// TODO 原因  值类型在反射中不能寻找存储在变量中具体值的地址, 所以值不能调用接收者为指针的方法
		//           指针类型在反射中调用接收者为值的方法时, 会自动解引用, 转为对应的值, 所以指针可以调用绑定的所有方法
		//           (不确定, 在学习接口时, 推断而来)
		for i := 0; i < v.NumMethod(); i++ {
			tm := v.Type().Method(i) // Method
			// vm := v.Method(i)       // Value
			fmt.Printf("  %-7v : %v\n", tm.Name, tm.Type)
		}
	}

	u := User{1, "name", 0}
	fmt.Println("  ---  receiver : instance  ---")
	f(u)

	fmt.Println()
	fmt.Println("  ---  receiver : pointer   ---")
	f(&u)
}

func TestStructCallMethod(t *testing.T) {
	// 传user指针, 不然拿不到接收者为指针的方法
	v := reflect.ValueOf(&User{Id: 0, Name: "0", uint8: 0})

	assert.Equal(t, 0, int(v.Elem().FieldByName("Id").Int()))
	assert.Equal(t, "0", v.Elem().FieldByName("Name").String())

	v.MethodByName("SetId").Call([]reflect.Value{reflect.ValueOf(1)})
	v.MethodByName("SetName").Call([]reflect.Value{reflect.ValueOf("1")})

	assert.Equal(t, 1, int(v.MethodByName("GetId").Call(nil)[0].Int()))
	assert.Equal(t, "1", v.MethodByName("GetName").Call(nil)[0].String())
}
