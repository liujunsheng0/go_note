package channel

import (
	"reflect"
	"time"
	"unsafe"
)

func Now() string {
	return time.Now().Format("15:04:05")
}

func NewMutex() Mutex {
	// 实现互斥锁 chan 的长度为1
	return make(Mutex, 1)
}

func NewFibonacciSequence() *FibonacciSequence {
	f := FibonacciSequence{}
	f.init = false
	return &f
}

// runtime/chan.go
// 利用反射查看channel是否关闭
// true: 关闭 false: 未关闭
func IsClosed(ch interface{}) bool {
	if reflect.TypeOf(ch).Kind() != reflect.Chan {
		panic("only channel!")
	}
	ptr := *(*uintptr)(unsafe.Pointer(unsafe.Pointer(uintptr(unsafe.Pointer(&ch)) + unsafe.Sizeof(uint(0)))))
	// see hchan on https://github.com/golang/go/blob/master/src/runtime/chan.go
	// type hchan struct {
	// qcount   uint           // total data in the queue
	// dataqsiz uint           // size of the circular queue
	// buf      unsafe.Pointer // points to an array of dataqsiz elements
	// elemsize uint16
	// closed   uint32         // 0未关闭, >0关闭
	// **

	ptr += unsafe.Sizeof(uint(0)) * 2
	ptr += unsafe.Sizeof(unsafe.Pointer(uintptr(0)))
	ptr += unsafe.Sizeof(uint16(0))
	return *(*uint32)(unsafe.Pointer(ptr)) > 0
}

// 从chan接收值
// Receive (接收的值, true(未关闭), 是否超时); (0, false(关闭), 是否超时)
func Receive(ch chan int) (int, bool, bool) {
	select {
	// v, ok := <- channel
	// 如果 channel 未关闭  阻塞   v为接收的值   ok等于true
	// 如果 channel 关闭   不阻塞  v为默认值     ok等于false
	case v, ok := <-ch:
		return v, ok, false
	case <-time.After(5 * time.Millisecond):
		return 0, true, true
	}
}
