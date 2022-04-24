package closure

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
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

// 闭包实现
// https://www.bilibili.com/video/BV1hv411x7we?p=7&spm_id_from=pageDriver
// https://www.bilibili.com/video/BV1hv411x7we?p=8
func TestClosure2(t *testing.T) {
	ar := assert.New(t)
	f := func() (fs []func()) {
		// 闭包导致逃逸, i分配到堆上
		for i := 0; i < 5; i++ {
			fs = append(fs, func() {
				ar.Equal(i, 5)
			})
		}
		return
	}
	fs := f()
	for i := 0; i < len(fs); i++ {
		fs[i]()
	}
}
