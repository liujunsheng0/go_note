package channel

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// Go会自动检查明显的死锁
// 通信是一种同步形式: 通过通道, 两个协程在通信中同步交换数据, 无缓冲通道成为了多个协程同步的完美工具
func TestAsync(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(2)
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		time.Sleep(time.Second)
		t.Log("send ch1 1", Now())
		ch1 <- 1
		t.Log("wait ch2  ", Now())
		t.Log("recv ch2", <-ch2, Now())
		wg.Done()
	}()
	go func() {
		t.Log("wait ch1  ", Now())
		t.Log("recv ch1", <-ch1, Now())
		time.Sleep(time.Second)
		t.Log("send ch2 2", Now())
		ch2 <- 2
		wg.Done()
	}()
	wg.Wait()
}

// 等待函数执行结果
func TestWaitResult(t *testing.T) {
	ch := make(chan int)
	sum := func(a []int, ch chan int) {
		sum := 0
		for _, i := range a {
			time.Sleep(1 * time.Second)
			sum += i
		}
		ch <- sum
	}
	go sum([]int{1, 2, 3}, ch)

	// do something

	t.Log("wait result", Now())
	t.Log(fmt.Sprintf("recv %-6v", <-ch), Now())
}

// channel实现互斥锁
func TestImplementMutex(t *testing.T) {
	fmt.Println("lock  unlock    time")
	work := func(i int, m Mutex, wg *sync.WaitGroup) {
		defer wg.Done()
		// 加锁
		m.Lock()
		fmt.Printf("%-4v  %-6v  %v\n", i, "", Now())
		// 解锁
		defer func() {
			m.UnLock()
			fmt.Printf("%-4v  %-6v  %v\n", "", i, Now())
		}()

		// do something
		time.Sleep(time.Duration(i) * time.Second)
	}

	delta := 5
	wg := sync.WaitGroup{}
	wg.Add(delta)
	m := NewMutex()
	for i := 1; i < delta+1; i++ {
		go work(i, m, &wg)
	}
	wg.Wait()
}

// channel实现类似于Python中的生成器
func TestGenerate(t *testing.T) {
	fib := NewFibonacciSequence()
	fmt.Println("idx  value")

	fmt.Println("---  next   ---")
	for i := 0; i < 3; i++ {
		fmt.Println(i, "    ", fib.Next())
	}

	fmt.Println("---  chan   ---")
	i := 1
	for v := range fib.Chan() {
		if i > 3 {
			break
		}
		i++
		fmt.Println(i, "    ", v)
	}

	fmt.Println("--- close  ---")
	for i := 0; i < 3 && !fib.IsClosed(); i++ {
		fmt.Println(i, "   ", fib.Next())
		fib.Close()
	}
}
