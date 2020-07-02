package _defer

import "testing"

// defer 允许在函数返回之前执行指定的语句
// 将要执行的操作放到栈里(逆序执行), 在return之前执行
func TestDefer(t *testing.T) {
	defer t.Log(1)
	defer t.Log(2)
	defer func() {
		for i := 5; i > 2; i-- {
			t.Log(i)
		}
	}()
	// defer_recover_test.go:13: 5
	// defer_recover_test.go:13: 4
	// defer_recover_test.go:13: 3
	// defer_recover_test.go:20: 2
	// defer_recover_test.go:20: 1
}

// panic, recover用于错误处理机制
func TestRecover(t *testing.T) {
	defer func() {
		// 类似于try - catch
		// 用于捕获函数执行过程中的异常, 适用范围: 仅在当前函数中
		if err := recover(); err != nil {
			t.Log(err)
		}
	}()
	panic("test raise exception")
}
