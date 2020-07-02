# [new() make()](https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/07.2.md)

看起来二者没有什么区别，都在堆上分配内存，但是它们的行为不同，适用于不同的类型。

- new(T) 为每个新的类型T分配一片内存，初始化为 0 并且返回类型为*T的内存地址。

  这种方法**返回一个指向类型为T，值为 0 的地址的指针**，它适用于值类型如int，数组，结构体；它相当于 `&T{}`。

- make(T) **返回一个类型为 T 的初始值**，它只适用于3种内建的引用类型：切片、map 和 channel

换言之，new 函数分配内存，make 函数初始化。见下图

![](new_make.png)

```go
var p1 *[]int = new([]int)        // *p1 == nil; with len and cap 0
var p2 []int  = make([]int, 0)    // 切片已经被初始化, 但是指向一个空的数组
fmt.Println(reflect.TypeOf(p1), reflect.TypeOf(p2)) // *[]int []int
```



## 如何理解new、make、slice、map、channel的关系

1. slice、map以及channel都是golang内建的一种引用类型，三者在内存中存在多个组成部分， 需要对内存组成部分初始化后才能使用，而make就是对三者进行初始化的一种操作方式
2. new 获取的是存储指定变量内存地址的一个变量，对于变量内部结构并不会执行响应的初始化操作， 所以slice、map、channel需要**make进行初始化并获取对应的内存地址，而非new简单的获取内存地址**

