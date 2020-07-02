package channel

// 斐波那契数列
// 斐波那契数列指的是这样一个数列 1,1,2,3,5,8,13,21,34,55,89,144,233,377,610,987,1597,2584,4181,6765,10946,17711 ......
// 这个数列从第3项开始, 每一项都等于前两项之和
// 实现类似于Python生成器
type FibonacciSequence struct {
	ch   chan int
	init bool // 是否已经初始化
}

// 生成器
// 初始化函数
func (s *FibonacciSequence) Init() {
	if s.init {
		return
	}
	s.ch = make(chan int, 1)
	s.init = true
	go func() {
		x := 1
		y := 1
		s.ch <- x
		s.ch <- y
		for !IsClosed(s.ch) {
			z := x + y
			s.ch <- z
			x, y = y, z
		}
	}()
}

func (s *FibonacciSequence) Chan() chan int {
	if !s.init {
		s.Init()
	}
	return s.ch
}

// 获取下一个元素
func (s *FibonacciSequence) Next() int {
	if !s.init {
		s.Init()
	}
	return <-s.ch
}

func (s *FibonacciSequence) IsClosed() bool {
	return IsClosed(s.ch)
}

func (s FibonacciSequence) Close() {
	close(s.ch)
	s.init = false
}
