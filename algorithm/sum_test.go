package algorithm

import (
	"sort"
	"testing"
)

// O(N*N)
func threeSum1(nums []int) [][]int {
	numMap := map[int]int{}
	for _, num := range nums {
		numMap[num] += 1
	}
	sort.Ints(nums)
	ans := make([][]int, 0, 5)
	for i, a := range nums {
		if i > 0 && a == nums[i-1] {
			continue
		}
		numMap[a] -= 1
		tmp := nums[i+1:]
		for j, b := range tmp {
			if (j > 0 && b == tmp[j-1]) || (-a-b) < b {
				continue
			}
			numMap[b] -= 1
			numMap[-a-b] -= 1
			if numMap[-a-b] >= 0 && numMap[a] >= 0 && numMap[b] >= 0 {
				ans = append(ans, []int{a, b, -a - b})
			}
			numMap[b] += 1
			numMap[-a-b] += 1
		}
		numMap[a] += 1
	}
	return ans
}

func threeSum2(nums []int) [][]int {
	sort.Ints(nums)
	ans := make([][]int, 0, 5)
	length := len(nums)
	for i := 0; i < length-1; i++ {
		num := nums[i]
		if num > 0 {
			break
		}
		if i > 0 && num == nums[i-1] {
			continue
		}
		target := -num
		left := i + 1
		right := length - 1
		for left < right {
			if nums[left]+nums[right] == target {
				ans = append(ans, []int{num, nums[left], nums[right]})
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if nums[left]+nums[right] > target {
				right--
			} else {
				left++
			}
		}
	}
	return ans
}

func Test3Sum(t *testing.T) {
	nums := []int{-2, 0, 1, 1, 2}
	t.Log(threeSum1(nums))
	t.Log(threeSum2(nums))
}
