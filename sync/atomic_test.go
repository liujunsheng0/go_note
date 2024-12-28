package sync_

import (
	"sync/atomic"
	"testing"
	"unsafe"
)

func TestSwapPointer(t *testing.T) {
	a := 1
	aPtr := unsafe.Pointer(&a)
	b := 2
	t.Log(a, b, aPtr, *((*int)(aPtr)))
	//
	// func CompareAndSwapPointer(addr *unsafe.Pointer, old, new unsafe.Pointer) (swapped bool)
	// 交互地址 if addr 与old地址相同, 则将new->addr return: 是否将new->addr
	t.Log(atomic.CompareAndSwapPointer(&aPtr, unsafe.Pointer(&a), unsafe.Pointer(&b)))
	t.Log(atomic.CompareAndSwapPointer(&aPtr, unsafe.Pointer(&a), unsafe.Pointer(&b)))
	t.Log(a, b, aPtr, *((*int)(aPtr)))
}
