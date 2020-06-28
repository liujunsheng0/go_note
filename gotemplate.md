# 包的分级声明和初始化

可以在导入包之后定义或声明常量、变量和类型，这些对象的作用域都是全局的，可以被本包中所有的函数调用，然后声明一个或多个函数和方法。



**Go 程序的执行（程序启动）顺序如下：**

1. 按顺序导入所有被 main 包引用的其它包，然后在每个包中执行如下流程；
2. 如果该包又导入了其它的包，则从第一步开始递归执行，但是每个包只会被导入一次；
3. 然后以相反的顺序在每个包中初始化常量和变量，如果该包含有 init 函数的话，则调用该函数，如果有多个init函数，则按顺序执行init函数；
4. 在完成这一切之后，main 也执行同样的过程，最后调用 main 函数开始执行程序；



# **命名规范**

1. 命名不需要指出自己所属的包，因为在调用的时候会使用包名作为限定符。

2. 返回某个对象的函数或方法的名称一般都是使用名词，没有Get... 之类的字符，如果是用于修改某个对象，则使用 SetName；
3. 有必须要的话可以使用大小写混合的方式，如 SchoolName 或 schoolName，而不是使用下划线来分割多个名称；



# 文件模板

```go
package main

import (
	"fmt"
)

const (
	c = ""
)

var (
	v int
)

type T struct{}

func init() {
	// initialization of package
}

func main() {
	var a int
	Func1()
	// ...
	fmt.Println(a)
}

func (t T) Method1() {
	//...
}
// method...

func Func1() { // exported function Func1
	//...
}
// func...
```