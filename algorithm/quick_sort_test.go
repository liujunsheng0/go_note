package algorithm

import (
	"fmt"
	"sort"
	"testing"
)

// 快排
func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	start, end := 0, len(arr)-1
	compare := arr[0]
	index := 0
	for start < end {
		// 比num小
		for ; start < end && arr[end] > compare; end-- {
		}
		arr[index] = arr[end]
		// 比num大
		for ; start < end && arr[start] <= compare; start++ {
		}
		arr[end] = arr[start]
		index = start
	}
	arr[index] = compare
	QuickSort(arr[:index])
	QuickSort(arr[index+1:])
	return arr
}

func TestQuickSort(t *testing.T) {
	fmt.Println(sort.IntsAreSorted(QuickSort([]int{5, 1, 1, 2, 0, 0})))
	fmt.Println(sort.IntsAreSorted(QuickSort([]int{5, 2, 3, 1})))
	fmt.Println(sort.IntsAreSorted(QuickSort([]int{3, -1})))
	fmt.Println(sort.IntsAreSorted(QuickSort([]int{3, 1, 2, 2})))
	fmt.Println(sort.IntsAreSorted(QuickSort([]int{3, 1, 2})))
	fmt.Println(sort.IntsAreSorted(QuickSort([]int{-3, -1, -2})))

}
