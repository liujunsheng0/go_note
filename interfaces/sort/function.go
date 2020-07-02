package sort

func Sort(d Sorter) {
	// 冒泡, 不停的选最大
	for i := 1; i < d.Len(); i++ {
		for j := 0; j < d.Len()-i; j++ {
			if d.Less(j+1, j) {
				// Swap的接收者要为引用
				d.Swap(j, j+1)
			}
		}
	}
}
