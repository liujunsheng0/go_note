package sort

import "testing"

func TestSort(t *testing.T) {
	i := IntSlice{9, 2, 5, 3}
	f := Float32Slice{7.0, 1.0, 4.0, 2.0}
	s := StringSlice{"a", "b", "c", "aaaa", "bbb", "cc"}
	Sort(i)
	Sort(f)
	Sort(s)
	t.Log("int   :", i)
	t.Log("float :", f)
	t.Log("string:", s)
}
