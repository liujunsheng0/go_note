# [异常处理](https://learnku.com/docs/the-way-to-go/chapter-description/3673)

Go中没有`try-catch`机制，但是有一套`defer-panic-and-recover机制`。

Go 的设计者觉得 `try/catch` 机制的使用太泛滥了，而且**从底层向更高的层级抛异常太耗费资源**。他们给 Go 设计的机制也可以 “捕捉” 异常，但是更轻量，并且只应该作为处理错误的最后的手段。

Go 有一个预先定义的 error 接口类型

```go
type error interface {
    Error() string
}
```



## [定义错误](https://learnku.com/docs/the-way-to-go/131-error-handling/3674)



## [运行时异常](https://learnku.com/docs/the-way-to-go/132-runtime-exception-and-panic/3675)

当发生像数组下标越界或类型断言失败这样的运行错误时，Go 运行时会触发*运行时 panic*，伴随着程序的崩溃抛出一个 `runtime.Error` 接口类型的值。这个错误值有个 `RuntimeError()` 方法用于区别普通错误。

```go
// runtime.Error
type Error interface {
	error
	RuntimeError()
}
```

`panic` 可以直接从代码初始化：当程序不能继续运行时，可以使用 `panic`函数产生一个中止程序的运行时错误。`panic` 接收一个做任意类型的参数，通常是字符串，在程序死亡时被打印出来。

```go
func main() {
	panic("fatal")
}
```

Go panicking：

在多层嵌套的函数调用中调用 panic，可以马上**中止当前函数的执行**，所有的 defer 语句都会保证执行并把控制权交还给接收到 panic 的函数调用者。这样向上冒泡直到最顶层，并执行（每层的） defer，在栈顶处程序崩溃，并在命令行中用传给 panic 的值报告错误情况：这个终止过程就是 *panicking*。

recover：

recover被用于从 panic 或 错误场景中恢复：让程序可以从 panicking 重新获得控制权，停止终止过程进而恢复正常执行。

`recover` 只能在 defer 修饰的函数中使用：用于取得 panic 调用中传递过来的错误值，如果是正常执行，调用 `recover` 会返回 nil，且没有其它效果。