package sort

// 接口的应用
// 实现package sort

type Sorter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

// 切片本身就是引用, 所以接收者是值就会影响调用者
type IntSlice []int
type Float32Slice []float32

// 以字符串长度排序, 如果长度相同, 按字典序排序
type StringSlice []string

// IntSlice
func (s IntSlice) Len() int           { return len(s) }
func (s IntSlice) Less(i, j int) bool { return s[i] < s[j] }
func (s IntSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

// Float32Slice
func (s Float32Slice) Len() int           { return len(s) }
func (s Float32Slice) Less(i, j int) bool { return s[i] < s[j] }
func (s Float32Slice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

// StringSlice
func (s StringSlice) Len() int      { return len(s) }
func (s StringSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s StringSlice) Less(i, j int) bool {
	if len(s[i]) != len(s[j]) {
		return len(s[i]) < len(s[j])
	}
	return s[i] < s[j]
}
