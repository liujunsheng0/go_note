package channel

import (
	"math/rand"
	"testing"
	"time"
)

// Ticker 以指定的时间间隔重复的向通道 Ticker.C 发送时间值, 用处: 如可以限制处理频率
func TestTicker(t *testing.T) {
	// 时间间隔
	timeInterval := time.Second
	ticker := time.NewTicker(timeInterval)
	// 调用 Stop() 使计时器停止
	defer ticker.Stop()

	for i := 0; i < 5; i++ {
		now := <-ticker.C
		t.Log(i, now.Format("15:04:05"))
	}
}

// Tick 返回一个通道而不必关闭它的时候这个函数非常有用, 它以 d 为周期给返回的通道发送时间
func TestTick(t *testing.T) {
	timeInterval := time.Second
	tick := time.Tick(timeInterval)
	for i := 0; i < 10; i++ {
		now := <-tick
		t.Log(now.Format("15:04:05"))
	}
	time.Sleep(time.Second * 2)
}

// 定时器, 可用于select
func TestAfter(t *testing.T) {
	t.Log("now  ", Now())
	ch := make(chan int, 1)
	defer close(ch)

	for i := 0; i < 3; i++ {
		go func() {
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2)+998+i))
			ch <- 1
		}()
		select {
		case <-ch:
			t.Log("ch   ", Now())
		case <-time.After(time.Second):
			t.Log("after", Now())
			return
		}
	}

}

// 简单超时模式 1s内没有收到数据, 则超时
func TestTimeOut(t *testing.T) {
	// 方法1
	// 注意缓冲大小设置为 1, 可以避免协程死锁以及确保超时的通道可以被垃圾回收
	timeout := make(chan bool, 1)
	ch := make(chan int, 1)
	go func() {
		time.Sleep(time.Second)
		timeout <- true
	}()
	t.Log("now     ", Now())
	select {
	case v := <-ch:
		t.Log("recv", v)
	case <-timeout:
		t.Log("timeout1", Now())
	}

	// 方法2
	select {
	case v := <-ch:
		t.Log("recv", v)
	case <-time.After(time.Second):
		t.Log("timeout2", Now())
	}

}
