package struct_

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

var (
	n Null
	p People
	s Student
)

func init() {
	s.IdCard = "1000"
	s.People.Name = "people"
	s.Name = "student"
	s.Age = 1
	s.SchoolId = "2020"
	s.School = "BeiJing"
	s.Grade = "5"
}

// 结构体初始化
func TestStructInit(t *testing.T) {
	var s1 Student
	s2 := Student{}
	// 有序初始化
	s3 := Student{People{IdCard: "1", Name: "n1", Age: 2}, "n2", "1", "school", "grade"}
	// 无序初始化
	s4 := Student{SchoolId: "1", School: "1", Grade: "2"}
	s5 := NewStudent("", "S", "G")
	assert.Equal(t, "id card=     name=     age=0  school id=     school=-", s1.String())
	assert.Equal(t, "id card=     name=     age=0  school id=     school=-", s2.String())
	assert.Equal(t, "id card=1    name=n2   age=2  school id=1    school=school-grade", s3.String())
	assert.Equal(t, "id card=     name=     age=0  school id=1    school=1-2", s4.String())
	assert.Equal(t, "id card=     name=     age=0  school id=     school=S-G", s5.String())
}

// 结构体指针初始化
func TestStructPointInit(t *testing.T) {
	// 直接定义指针为nil, 此时不能对p进行操作
	var p1 *Student
	// new可以分配内存来初始化结构休, 并返回分配的内存指针, 因为已经初始化了, 所以可以直接访问字段
	p2 := new(Student)

	assert.Equal(t, "*struct_.Student", reflect.TypeOf(p1).String())
	assert.Equal(t, "*struct_.Student", reflect.TypeOf(p2).String())

	assert.Nil(t, p1)
	assert.NotNil(t, p2)

}

// 接收者为 "结构体值" 和 "结构体指针" 的区别
// 接收者只要为 结构体值 时, 不影响原变量  (即调用者为结构体指针和结构体值时, 被调方法中使用的是结构体 *副本*)
// 接收者只要为 结构体指针 时, 会影响原变量  (即调用者为结构体指针和结构体值时, 被调方法中使用的是结构体 *引用*)
func TestStructReceiver(t *testing.T) {
	s.Age = 1
	s.SetAge(2) // 副本, 并不会影响变量s
	assert.Equal(t, 1, s.Age)

	(&s).SetAge(3) // 副本, 指针调用了接收者为结构体值方法时, 使用的仍然是副本, 并不会影响变量s
	assert.Equal(t, 1, s.Age)

	s.SetAgeReference(4) // 引用, 会影响变量s
	assert.Equal(t, 4, s.Age)
}

// GetSize 为获取结构体中Null值的大小
func TestStructSize(t *testing.T) {
	// Null的size
	assert.Equal(t, 0, n.GetSize())
	assert.Equal(t, 0, p.GetSize())
	assert.Equal(t, 0, s.GetSize())

	// 对应结构体的size
	assert.Equal(t, 0, int(unsafe.Sizeof(n)))
	assert.Equal(t, 40, int(unsafe.Sizeof(p)))
	assert.Equal(t, 104, int(unsafe.Sizeof(s)))
}

// 优先访问外层字段, 外层名字会覆盖内层名字, 但是两者的内存空间都保留
// 结构体匿名字段
func TestStructAnonymousField(t *testing.T) {
	// 访问匿名字段中被覆盖的属性
	assert.Equal(t, "people", s.People.Name)
	assert.Equal(t, "student", s.Name)
}

// 接收者为 "结构体值" 和 "结构体指针" 时, 利用反射获取的属性不同
// 结构体值: reflect 能获取字段    仅能获取接收者为值的方法
// 结构体指针: reflect 不能获取字段   能获取所有方法
func TestStructReflect(t *testing.T) {
	f := func(i interface{}) {
		v := reflect.TypeOf(i)
		// 仅适用于结构体值
		if v.Kind() == reflect.Struct {
			for i := 0; i < v.NumField(); i++ {
				m := v.Field(i)
				t.Log(fmt.Sprintf("%-8v: %v", m.Name, m.Type))
			}
		}
		for i := 0; i < v.NumMethod(); i++ {
			m := v.Method(i)
			t.Log(fmt.Sprintf("%-8v: %v", m.Name, m.Type))
		}
	}

	t.Log("---  receiver : instance  ---")
	f(s)
	t.Log()
	t.Log("---  receiver : pointer   ---")
	f(&s)
}
