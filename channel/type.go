package channel

// 信号
type Mutex chan interface{}

// 互斥锁 - 加锁
func (m Mutex) Lock() {
	if cap(m) != 1 {
		panic("cap(channel) != 1")
	}
	m <- 1
}

// 互斥锁- 解锁
func (m Mutex) UnLock() {
	if len(m) != 1 {
		panic("not lock")
	}
	<-m
}
