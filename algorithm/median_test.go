package algorithm

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 中位数

// 给定一个未排序的整数数组，找到其中位数。中位数是排序后数组的中间值，如果数组的个数是偶数个，则返回排序后数组的第N/2个数。
// https://www.lintcode.com/problem/80/
// 数组大小不超过10000
//输入: [4, 5, 1, 2, 3]  输出: 3
func medianArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	if len(nums)%2 == 0 {
		return nums[len(nums)/2-1]
	}
	return nums[len(nums)/2]
}

func TestMedianArray(t *testing.T) {
	ar := assert.New(t)
	ar.Equal(3, medianArray([]int{4, 5, 1, 2, 3}))
}

// 两个排序数组的中位数
// 两个排序的数组A和B分别含有m和n个数，找到两个排序数组的中位数，要求时间复杂度应为O(log (m+n))
// https://www.lintcode.com/problem/65/solution
// 中位数的定义: 中位数是排序后数组的中间值
// 		如果有数组中有n个数且n是奇数，则中位数为A[(n−1)/2]。
// 		如果有数组中有n个数且n是偶数，则中位数为 (A[(n−1)/2]+A[(n−1)/2+1])/2.
// 比如: 数组A=[1,2,3]的中位数是2，数组A=[1,19]的中位数是10。
// k: 第k个数
func medianSortedArrays(A []int, B []int, k int) int {
	if k < 1 {
		return 0
	}
	if len(A) == 0 {
		return B[k-1]
	}
	if len(B) == 0 {
		return A[k-1]
	}
	if k == 1 {
		return MinInt(A[0], B[0])
	}
	aK := MinInt(k/2, len(A)) - 1
	bK := MinInt(k/2, len(B)) - 1
	if A[aK] < B[bK] {
		return medianSortedArrays(A[aK+1:], B, k-aK-1)
	}
	return medianSortedArrays(A, B[bK+1:], k-bK-1)
}

func TestMedianSortedArrays(t *testing.T) {
	ar := assert.New(t)
	for i := 0; i < 100; i++ {
		a := []int{0}
		b := []int{0}
		for j := 0; j < rand.Intn(100); j++ {
			if rand.Intn(2) == 0 {
				a = append(a, a[len(a)-1]+rand.Intn(10))
			}
			if rand.Intn(2) == 0 {
				b = append(b, b[len(b)-1]+rand.Intn(10))
			}
		}
		sortArray := MergeAscArray(a, b)
		k := rand.Intn(len(a)+len(b)) + 1
		ar.Equal(sortArray[k-1], medianSortedArrays(a, b, k))
	}
}
