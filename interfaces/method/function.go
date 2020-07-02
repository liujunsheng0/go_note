package method

func AppendNums(a Appender, start, end int) {
	for i := start; i <= end; i++ {
		a.Append(i)
	}
}

func IsEmpty(l Lener) bool {
	return l.Len() == 0
}
