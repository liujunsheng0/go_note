package algorithm

import (
	"math/rand"
	"sort"
	"testing"
)

// 堆排序
// 子结点的键值或索引总是小于（或者大于）它的父节点
func AdjustHeap(arr []int, root int) {
	if len(arr) < 2 {
		return
	}
	left := 2*root + 1
	right := 2*root + 2
	if left < len(arr) && arr[left] > arr[root] {
		arr[left], arr[root] = arr[root], arr[left]
		AdjustHeap(arr, left)
	}
	if right < len(arr) && arr[right] > arr[root] {
		arr[right], arr[root] = arr[root], arr[right]
		AdjustHeap(arr, right)
	}
}

func HeapSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	// 调整堆
	for i := (len(arr) + 1) / 2; i >= 0; i-- {
		AdjustHeap(arr, i)
	}
	for i := len(arr) - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		AdjustHeap(arr[:i], 0)
	}
	return arr
}

func TestHeapSort(t *testing.T) {
	for i := 0; i < 100; i++ {
		arr := []int{3, 4, 5, 2, 1, 3, -1, -1, -1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, -2, 0, -100, -21, -30, -50}
		for i := 0; i < 100; i++ {
			arr = append(arr, rand.Intn(1000)-400)
		}
		arr1 := make([]int, 0, len(arr))
		HeapSort(arr)
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
