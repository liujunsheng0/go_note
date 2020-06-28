# [Go - unsafe](https://studygolang.com/pkgdoc)

## Poniter

```go
// 表示任意类型
type ArbitraryType int
type Pointer *ArbitraryType
```

Pointer类型用于表示任意类型的指针，类似于c++中的void*，可进行如下操作

1.  任意类型的指针   可以转换为     Pointer类型
2. Pointer类型         可以转换为     任意类型的指针
3. uintptr类型          可以转换为     Pointer类型值
4. Pointer类型          可以转换为     uintptr类型值

> **Pointer类型允许程序绕过类型系统读写任意内存，使用它时必须谨慎**



## Sizeof

```go
func Sizeof(v ArbitraryType) uintptr
```

返回v本身数据所占用的字节数，该大小不包括v引用的内存。

例如

+ 若v是一个切片，它会返回该切片描述符的大小，而非该切片底层引用的内存大小
+ 若v是字符串，字符串对应结构体中的指针和字符串长度部分，并不包含指针指向的字符串的内容，其他类型也是如此，只包含结构体中定义的变量，并不包含其引用的内容



## Alignof

```go
func Alignof(v ArbitraryType) uintptr
```

返回v的对齐方式（即类型v在内存中占用的字节数）， 若v是结构体类型的字段的形式，它会返回字段在该结构体中的对齐方式。



## Offsetof

```go
// v必须为结构体类型
func Offsetof(v ArbitraryType) uintptr
```


返回该结构起始处与该字段起始处之间的字节数



## uintptr

```go
// uintptr is an integer type that is large enough to hold the bit pattern of
// any pointer.
type uintptr uintptr
```

内置类型，一个整数类型，其大小足以容纳任何指针的位模式



## 指针

Go中的指针

1. 原生指针

   原生指针(&变量)不支持运算，即不能对原生指针做加减法等运算

2.  uintptr

   **用来进行指针计算**，因为它是整型，所以很容易计算出下一个指针所指向的位置

3. unsafe.Pointer

   可指向任意类型的指针，类似于中间件，作为不同类型指针互相转换的桥梁

   不支持指针运算

### 指针定义和初始化

```go
var pointName *变量类型  // 默认值为nil, 未分配内存
pointName := &变量
pointName := new(int)   // 分配足够的内存以容纳该类型的值, 并返回指向它的指针
```

> 当指针为nil时，不能对(*pointName)进行操作包括读取，否则会报空指针异常

### 指针运算

```go
var v int
p1 := unsafe.Pointer(&v)    // *int    ->  Pointer
p2 := uintptr(p1)           // Pointer ->  uintptr
进行指针操作                  // uintptr +/- uintptr
p3 := unsafe.Pointer(p2)    // uintptr ->  Pointer
pInt := (*int)(p2)          // Pointer ->  *int
```

### 返回值

返回局部变量的指针是安全的，编译器会根据需要将其分配在 GC Heap或栈 上

