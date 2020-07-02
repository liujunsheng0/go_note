package method

// 接口与方法集

import (
	"fmt"
	"testing"
)

// 在 list  (值)   上调用 AppendNums  时会导致一个编译器错误, 因为 AppendNums 需要一个 Appender, 而它的方法 Append 的接收者为指针
// 在 list  (值)   上调用 IsEmpty     是可以的, 因为Len的接收者为实例(值)
// 在 &list (指针) 上调用 AppendNums  是可以的, 因为 AppendNums 需要一个 Appender, 并且它的方法 Append 的接收者为指针
// 在 &list (指针) 上调用 IsEmpty     是可以的, 因为指针会被自动解引用, 转化为值
func TestList(t *testing.T) {
	f := func(l Lener) {
		if IsEmpty(l) {
			t.Log(fmt.Sprintf("%-15T list length is     empty", l))
		} else {
			t.Log(fmt.Sprintf("%-15T list length is not empty", l))
		}
	}
	var list List
	// compiler error:
	// cannot use lst (type List) as type Appender in argument to AppendNums:
	//       List does not implement Appender (AppendNums method has pointer receiver)
	// 接口变量, 接收者为指针, 值无法调用接收者为指针的方法
	// AppendNums(list, 1, 10)
	f(list)

	AppendNums(&list, 0, 10)
	f(&list)
}
