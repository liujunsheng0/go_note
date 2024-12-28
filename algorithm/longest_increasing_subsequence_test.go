package algorithm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
https://leetcode-cn.com/problems/longest-increasing-subsequence/
最长递增子序列
给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。
子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。
*/
func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	ans := 0
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			// 以nums[i]结尾的最长递增子序列
			if nums[i] > nums[j] {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
				}
			}
		}
		ans = MaxInt(ans, dp[i])
	}
	return ans
}

func TestLengthOfLIS(t *testing.T) {
	//assert.Equal(t, 4, lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	// ans: 2, 3, 7, 18, 19, 20, 21
	// nums: 10, 9, 2, 5, 3, 7, 101, 18, 19, 20, 21, 2, 12
	// dp  : 1   1  1  2  2  3   4    4   5   6   7  7  7
	assert.Equal(t, 7, lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18, 19, 20, 21, 2, 12}))
}
