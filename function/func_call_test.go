package function

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func incr1(a int) int {
	var b int
	defer func() {
		a++
		b++
	}()
	a++
	b = a
	return b
}

func incr2(a int) (b int) {
	defer func() {
		a++
		b++
	}()
	a++
	return a
}

/*
https://www.bilibili.com/video/BV1hv411x7we?p=6
go 语言的函数栈帧是一次性分配
bp: 栈底（栈基）
sp: 栈顶（栈指针）
Test
incr1 栈帧布局
	1) a = 0
	2) b = 0
	3) 0     incr1的返回值 (return value)
    4) 0     (函数参数a: 函数的参数空间, 值拷贝, 从右-左)
	5) return addr  incr1的返回地址
    6) main的bp地址
	7) b = 0   局部变量
	8) return 编译器插入指令, 恢复到调用者栈
           执行顺序: 先给返回值赋值(即将b的值拷贝到return value处), 在执行defer函数, 在恢复到调用者栈
	过程:
	a++, 4) a=1
    b=a, 7) b=1
    return b:  (1): 将b的值拷贝到3)
               (2): 执行defer, 4) a=2  7) b=2
               (3): 释放栈空间, 恢复到调用者栈
incr2 栈帧布局
	1) a = 0
	2) b = 0
	3) 0     incr2的返回值b
    4) 0     (函数参数a: 函数的参数空间, 值拷贝, 从右-左)
	5) return addr  incr1的返回地址
    6) main的bp地址
	7) return
	过程:
	a++, 4) a=1
    return b:  (1): 将a的值拷贝到3), 此事b=1
               (2): 执行defer, 4) a=2  7) b=2
               (3): 释放栈空间, 恢复到调用者栈
*/
func TestIncr(t *testing.T) {
	var a, b int
	ar := assert.New(t)
	b = incr1(a)
	ar.Equal(a, 0)
	ar.Equal(b, 1)
	b = 0
	b = incr2(a)

	ar.Equal(a, 0)
	ar.Equal(b, 2)
}
