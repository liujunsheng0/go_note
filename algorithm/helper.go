package algorithm

func MaxInt(arr ...int) int {
	if len(arr) == 0 {
		return 0
	}
	m := arr[0]
	for _, i := range arr[1:] {
		if i > m {
			m = i
		}
	}
	return m
}

func MinInt(arr ...int) int {
	if len(arr) == 0 {
		return 0
	}
	m := arr[0]
	for _, i := range arr[1:] {
		if i < m {
			m = i
		}
	}
	return m
}

func MaxFloat64(arr ...float64) float64 {
	if len(arr) == 0 {
		return 0
	}
	m := arr[0]
	for _, i := range arr[1:] {
		if i > m {
			m = i
		}
	}
	return m
}

// MergeAscArray 合并升序数组
func MergeAscArray(a, b []int) []int {
	ret := make([]int, 0, len(a)+len(b))
	ida := 0
	idb := 0
	for ida < len(a) && idb < len(b) {
		if a[ida] < b[idb] {
			ret = append(ret, a[ida])
			ida++
		} else {
			ret = append(ret, b[idb])
			idb++
		}
	}
	for ; ida < len(a); ida++ {
		ret = append(ret, a[ida])
	}
	for ; idb < len(b); idb++ {
		ret = append(ret, b[idb])
	}
	return ret
}
