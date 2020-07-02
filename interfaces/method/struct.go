package method

type Appender interface {
	Append(int)
}

type Lener interface {
	Len() int
}

type List []int

// 接收者为实例
func (l List) Len() int {
	return len(l)
}

// 接收者为指针
func (l *List) Append(val int) {
	*l = append(*l, val)
}
