package closure

import (
	"sync"
	"testing"
)

// 闭包 - 实现装饰器
// 利用闭包计算函数执行时间 (Python装饰器)
func TestClosure(t *testing.T) {
	var wg sync.WaitGroup
	var f func(...int) int

	wg.Add(6)

	// 和
	f = CalculateRunTime(Sum, &wg, t)
	go f(2)
	go f(2, 4)

	// 积
	f = CalculateRunTime(Multi, &wg, t)
	go f(2)
	go f(2, 4)

	// 幂
	f = CalculateRunTime(Pow, &wg, t)
	f(2, 3)
	f(2, 4)

	wg.Wait()
}
