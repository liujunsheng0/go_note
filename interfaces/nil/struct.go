package nil

import (
	"unsafe"
)

type Animal interface {
	Name() string
}

type Dog struct {
}

func (d *Dog) Name() string {
	if d == nil {
		return ""
	}
	return "dog"
}

// 在Go语言中, 一个interface{}类型的变量包含了2个指针, 一个指针指向值的类型, 另外一个指针指向实际的值
// InterfaceStructure 定义了一个interface{}的内部结构
type InterfaceStructure struct {
	pt uintptr // 到值类型的指针
	pv uintptr // 到值内容的指针
}

// asInterfaceStructure 将一个interface{}转换为InterfaceStructure
func asInterfaceStructure(i interface{}) InterfaceStructure {
	return *(*InterfaceStructure)(unsafe.Pointer(&i))
}
