package unsafe

type A struct {
	a int
}

// A放在首位
type B struct {
	A
	b int
}

// A放在后面
type C struct {
	c int
	A
}

func setA(p *A, i int) {
	p.a = i
}

func (s *A) setA(i int) {
	s.a = i
}
