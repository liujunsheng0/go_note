package algorithm

import (
	"math/rand"
	"sort"
	"testing"
)

// 快排
func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	start, end := 1, len(arr)-1
	compare := arr[0]
	index := start
	for start < end {
		// 比num大
		for ; start < end && arr[start] < compare; start++ {
		}
		arr[index] = arr[start]
		// 比num小
		for ; start < end && arr[end] >= compare; end-- {
		}
		arr[start] = arr[end]
		index = end
	}
	arr[index] = compare

	QuickSort(arr[0:index])
	QuickSort(arr[index+1:])
	return arr
}

func TestQuickSort(t *testing.T) {
	for i := 0; i < 100; i++ {
		arr := []int{3, 4, 5, 2, 1, 3, -1, -1, -1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, -2, 0, -100, -21, -30, -50}
		for i := 0; i < 100; i++ {
			arr = append(arr, rand.Intn(100)+-30)
		}
		arr1 := make([]int, 0, len(arr))
		QuickSort(arr)
		for i := range arr {
			arr1 = append(arr1, arr[i])
		}
		sort.Ints(arr)

		for i := range arr {
			if arr[i] != arr1[i] {
				t.Log("not equal", arr[i], arr1[i])
			}
		}
	}

}
