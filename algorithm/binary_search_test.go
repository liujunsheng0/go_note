package algorithm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 恢复旋转排序数组
// https://www.lintcode.com/problem/39/
// 给定一个旋转排序数组，在原地恢复其排序。（升序） 什么是旋转数组？
// 比如，原始数组为[1,2,3,4], 则其旋转数组可以是[1,2,3,4], [2,3,4,1], [3,4,1,2], [4,1,2,3]
// 输入：数组 = [4,5,1,2,3]  输出：[1,2,3,4,5]
// 使用O(1)的额外空间和O(n)时间复杂度
func RecoverRotatedSortedArray(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	n := nums[0]
	idx := 0
	for i, num := range nums {
		if num < n {
			idx = i
			break
		}
	}
	tmp := nums[idx:]
	tmp = append(tmp, nums[:idx]...)
	return tmp
}

func TestRecoverRotatedSortedArray(t *testing.T) {
	ar := assert.New(t)
	arr := []int{1, 2, 3, 5}
	ar.Equal(arr, RecoverRotatedSortedArray([]int{4, 1, 2, 3}))
	ar.Equal(arr, RecoverRotatedSortedArray([]int{3, 4, 1, 2}))
	ar.Equal(arr, RecoverRotatedSortedArray([]int{2, 3, 4, 1}))
}
