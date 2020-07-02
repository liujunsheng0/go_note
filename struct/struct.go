package struct_

import (
	"fmt"
	"unsafe"
)

// 空结构 节省 内存, 如用来实现 set 数据结构, set := make(map[string]struct{})
type Null struct{}

type People struct {
	Null
	IdCard string
	Name   string
	Age    int
}

type Student struct {
	People          // 匿名字段
	Name     string // 命名冲突, 外层名字会覆盖People中的Name, 但是两者的内存空间都保留
	SchoolId string
	School   string
	Grade    string
}

// 工厂方式实现结构体构造函数
// 可以将结构体名字小写, 使其成为私有, 实例化将通过工厂函数
func NewStudent(name, school, grade string) *Student {
	s := new(Student)
	s.Name = name
	s.School = school
	s.Grade = grade
	return s
}

// ps:当struct作为匿名字段出现时, struct.GetSize()等于 struct.Null.GetSize()
func (s Null) GetSize() int {
	return int(unsafe.Sizeof(s))
}

// 接收者是值, 传递的是变量副本, 不会影响原变量
func (s People) SetAge(age int) {
	s.Age = age
}

// 接收者是指针, 会影响原变量
func (s *People) SetAgeReference(age int) {
	s.Age = age
}

func (s *People) String() string {
	return fmt.Sprintf("id card=%-4s name=%-4s age=%-2d", s.IdCard, s.Name, s.Age)
}

// 接收者为值:   fmt.Println打印值/指针时,  会自动调用其String方法
// 接收者为指针: fmt.Println打印    指针时, 会自动调用其String方法(值不会)
// 因为实现原理为接口, 详见接口
// 当接收者为值时    变量为值或指针赋给接口变量时, 能调用此方法
// 当接收者为指针时  变量为指针赋给接口变量时,     能能调用此方法
// 当接收者为指针时  变量为值赋给接口变量时,      不能能调用此方法

func (s Student) String() string {
	return fmt.Sprintf("id card=%-4s name=%-4s age=%-2d school id=%-4s school=%s-%s",
		s.IdCard, s.Name, s.Age, s.SchoolId, s.School, s.Grade)
}
