package algorithm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 背包

// Backpack1 在n个物品中挑选若干物品装入背包，最多能装多满？假设背包的大小为m，每个物品的大小为Ai
// https://www.lintcode.com/problem/92/
func Backpack1(m int, A []int) int {
	if m < 1 || len(A) < 1 {
		return 0
	}
	dp := make([]int, m+1)
	// 保证每个物品只选一次, 防止重复计算
	for _, a := range A {
		// *** 由大->小填充 ***
		for i := m; i >= a; i-- {
			dp[i] = MaxInt(dp[i], dp[i-a], dp[i-a]+a)
		}
	}
	return dp[m]
}

func TestBackpack1(t *testing.T) {
	ar := assert.New(t)
	ar.Equal(9, Backpack1(10, []int{1, 5, 8, 3}))
	ar.Equal(11, Backpack1(11, []int{1, 5, 8, 3}))
}

// Backpack2 有 n 个物品和一个大小为 m 的背包. 给定数组 A 表示每个物品的大小和数组 V 表示每个物品的价值
// 问最多能装入背包的总价值是多大?
// https://www.lintcode.com/problem/125
func Backpack2(m int, A []int, V []int) int {
	if m < 1 || len(A) == 0 || len(A) != len(V) {
		return 0
	}
	dp := make([]int, m+1)
	for i, a := range A {
		for j := m; j >= a; j-- {
			dp[j] = MaxInt(dp[j], dp[j-a]+V[i], dp[j-a])
		}
	}
	return dp[m]
}

func TestBackpack2(t *testing.T) {
	ar := assert.New(t)
	ar.Equal(9, Backpack2(10, []int{2, 3, 5, 7}, []int{1, 5, 2, 4}))
	ar.Equal(10, Backpack2(10, []int{2, 3, 8}, []int{2, 5, 8}))
}

// 给定 n 种物品, 每种物品都有无限个. 第 i 个物品的体积为 A[i], 价值为 V[i].
// 再给定一个容量为 m 的背包. 问可以装入背包的最大价值是多少?
// https://www.lintcode.com/problem/440
func Backpack3(A []int, V []int, m int) int {
	if m < 1 || len(A) == 0 || len(A) != len(V) {
		return 0
	}
	dp := make([]int, m+1)
	// 不限次数
	for i := range dp {
		for j, a := range A {
			if a > i {
				continue
			}
			dp[i] = MaxInt(dp[i], dp[i-a]+V[j])
		}
	}
	return dp[m]
}

func TestBackpack3(t *testing.T) {
	ar := assert.New(t)
	ar.Equal(15, Backpack3([]int{2, 3, 5, 7}, []int{1, 5, 2, 4}, 10))
	ar.Equal(5, Backpack3([]int{1, 2, 3}, []int{1, 2, 3}, 5))
}

// 给出 n 个物品, 以及一个数组, nums[i]代表第i个物品的大小, 保证大小均为正数并且没有重复, 正整数 target 表示背包的大小, 找到能填满背包的方案数。
// 每一个物品可以使用无数次
// https://www.lintcode.com/problem/562
func Backpack4(nums []int, target int) int {
	if len(nums) == 0 || target < 1 {
		return 0
	}
	dp := make([]int, target+1)
	dp[0] = 1
	// 状态转移方程为 dp[j] = dp[j] + dp[j - nums[i]]
	for _, num := range nums {
		for i := num; i <= target; i++ {
			dp[i] += dp[i-num]
		}
	}
	return dp[target]
}

func TestBackpack4(t *testing.T) {
	ar := assert.New(t)
	ar.Equal(2, Backpack4([]int{2, 3, 6, 7}, 7))
}

// 给出 n 个物品, 以及一个数组, nums[i] 代表第i个物品的大小, 保证大小均为正数, 正整数 target 表示背包的大小, 找到能填满背包的方案数。
// 每一个物品只能使用一次
// https://www.lintcode.com/problem/563/
func Backpack5(nums []int, target int) int {
	if len(nums) == 0 || target < 1 {
		return 0
	}
	dp := make([]int, target+1)
	dp[0] = 1
	for _, num := range nums {
		for i := target; i >= num; i-- {
			dp[i] += dp[i-num]
		}
	}
	return dp[target]
}

func TestBackpack5(t *testing.T) {
	ar := assert.New(t)
	ar.Equal(2, Backpack5([]int{1, 2, 3, 3, 7}, 7))
}

// 你总共有n万元，希望申请国外的大学，要申请的话需要交一定的申请费用，给出每个大学的申请费用以及你得到这个大学offer的成功概率，大学的数量是 m。如果经济条件允许，你可以申请多所大学。找到获得至少一份工作的最高可能性。
// 0<=n<=10000,0<=m<=10000
// 输入:
//		n = 10
//		prices = [4, 4, 5]
//		probability = [0.1, 0.2, 0.3]
//	输出: 1 - 都拿不到的概率 = 获得至少一份工作的最高可能性
//       1 - (0.8 * 0.7)  = 0.440
//	解释：
//	选择第2和第3个学校
// https://www.lintcode.com/problem/800
func Backpack6(n int, prices []int, probability []float64) float64 {
	if n < 1 || len(prices) == 0 || len(prices) != len(probability) {
		return 0
	}
	// 获得至少一份工作的最高可能性的可能性
	dp := make([]float64, n+1)
	for i, price := range prices {
		for j := n; j >= price; j-- {
			dp[j] = MaxFloat64(dp[j], dp[j-price], 1-(1-dp[j-price])*(1-probability[i]))
		}
	}
	return dp[n]
}

func TestBackpack6(t *testing.T) {
	ar := assert.New(t)
	ar.Equal(44, int(Backpack6(10, []int{4, 4, 5}, []float64{0.1, 0.2, 0.3})*100))
	ar.Equal(37, int(Backpack6(10, []int{4, 5, 6}, []float64{0.1, 0.2, 0.3})*100))
}

// 假设你身上有 n 元，超市里有多种大米可以选择，每种大米都是袋装的，必须整袋购买，给出每种大米的重量，价格以及数量，求最多能买多少公斤的大米
// https://www.lintcode.com/problem/798/
// 输入:  n = 8, prices = [3,2], weights = [300,160], amounts = [1,6]
// 输出:  640公斤
// 解释:  全买价格为2的米。
// 输入:  n = 8, prices  = [2,4], weight = [100,100], amounts = [4,2 ]
// 输出:  400
// 解释:  全买价格为2的米
func Backpack7(n int, prices []int, weights []int, amounts []int) int {
	if n < 1 || len(prices) == 0 || len(prices) != len(weights) || len(prices) != len(amounts) {
		return 0
	}
	// 最大重量
	dp := make([]int, n+1)

	for i, price := range prices {
		weight := weights[i]
		// 个数限制
		for j := 0; j < amounts[i]; j++ {
			for k := n; k >= price; k-- {
				dp[k] = MaxInt(dp[k], dp[k-price]+weight)
			}
		}
	}
	return dp[n]
}

func TestBackpack7(t *testing.T) {
	ar := assert.New(t)
	ar.Equal(640, Backpack7(8, []int{3, 2}, []int{300, 160}, []int{1, 6}))
	ar.Equal(400, Backpack7(8, []int{2, 4}, []int{100, 100}, []int{4, 2}))
	ar.Equal(1899, Backpack7(54, []int{10, 8, 8, 3, 13, 5, 10, 6, 11, 7, 3, 20, 15, 14, 20, 9, 7, 16, 13, 15}, []int{78, 74, 8, 56, 91, 177, 159, 66, 62, 143, 102, 195, 27, 199, 141, 21, 106, 122, 147, 76}, []int{3, 19, 3, 15, 2, 14, 5, 16, 10, 11, 10, 15, 10, 1, 18, 16, 2, 1, 14, 15}))
}

// 给一些不同价值和数量的硬币。找出[1，n]范围内的总值有多少种形成方式？
// https://www.lintcode.com/problem/799
// 输入: n = 5  value = [1,4]  amount = [2,1]
// 输出:  4
// 解释: 可以组合出1，2，4，5
func Backpack8(n int, values []int, amounts []int) int {
	if len(values) == 0 || len(values) != len(amounts) || n <= 0 {
		return 0
	}
	dp := make([]int, n+1)
	ret := 0
	for i, val := range values {
		for num := 1; num <= amounts[i]; num++ {
			for j := n; j >= val*num; j-- {
				if (j-val == 0 || dp[j-val] > 0) && dp[j] == 0 {
					dp[j] = 1
					ret++
				}
			}
		}
	}
	return ret
}

func TestBackpack8(t *testing.T) {
	ar := assert.New(t)
	ar.Equal(4, Backpack8(5, []int{1, 4}, []int{2, 1}))
	ar.Equal(8, Backpack8(10, []int{1, 2, 4}, []int{2, 1, 1}))
	ar.Equal(61097, Backpack8(61148,
		[]int{933, 559, 232, 352, 601, 244, 576, 10, 167, 634, 674, 348, 157, 363, 13, 339, 796, 65, 280, 795, 912, 848, 110, 137, 656, 40, 793, 914, 610, 944, 872, 153, 81, 902, 390, 127, 914, 991, 691, 764},
		[]int{488, 92, 194, 785, 382, 956, 786, 248, 375, 246, 677, 427, 898, 759, 214, 86, 417, 676, 463, 480, 999, 684, 105, 674, 174, 225, 338, 524, 861, 804, 970, 836, 808, 330, 496, 285, 569, 751, 158, 384}))
	ar.Equal(74935, Backpack8(74993,
		[]int{944, 830, 489, 956, 902, 144, 347, 861, 336, 127, 7, 705, 855, 432, 147, 485, 820, 571, 676, 364, 473, 462, 156, 861, 768, 885, 898, 996, 677, 999, 23, 974, 671, 14, 760, 689, 929, 942, 603, 241, 135, 824, 621, 159, 652, 885, 791, 693, 350, 634, 957, 681, 302, 700, 352, 732, 718, 983, 203, 997, 913, 44, 841, 67, 647, 87, 621, 164, 169, 666, 891, 387, 233, 958, 35, 462, 972, 943, 230, 174},
		[]int{395, 913, 780, 323, 123, 33, 608, 377, 254, 804, 797, 824, 604, 73, 343, 628, 404, 717, 275, 684, 225, 236, 94, 120, 481, 613, 307, 857, 754, 324, 314, 944, 590, 647, 579, 661, 800, 499, 560, 47, 116, 22, 602, 568, 17, 709, 993, 443, 872, 381, 415, 130, 23, 608, 217, 780, 864, 569, 955, 525, 962, 435, 546, 769, 598, 440, 505, 756, 22, 817, 255, 697, 499, 706, 516, 960, 121, 599, 658, 72}))
}

// 给一个 n 英寸长的杆子和一个包含所有小于 n 的尺寸的价格. 确定通过切割杆并销售碎片可获得的最大值.
// https://www.lintcode.com/problem/700
// 输入： [1, 5, 8, 9, 10, 17, 17, 20]  8
// 输出：22
// 解释：
// 长度    | 1   2   3   4   5   6   7   8
// ---------------------------------------
// 价格    | 1   5   8   9  10  17  17  20
// 切成长度为 2 和 6 的两段。

// 输入： [3, 5, 8, 9, 10, 17, 17, 20] 8
// 输出：24
// 解释：
// 长度    | 1   2   3   4   5   6   7   8
// ---------------------------------------
// 价格    | 3   5   8   9  10  17  17  20
// 切成长度为 1 的 8 段。

func Backpack9(prices []int, n int) int {
	if len(prices) == 0 || n == 0 {
		return 0
	}
	dp := make([]int, n+1)
	for i, price := range prices {
		length := i + 1
		for j := length; j <= n; j++ {
			dp[j] = MaxInt(dp[j], dp[j-length]+price)
		}
	}
	return dp[n]
}

func TestBackpack9(t *testing.T) {
	ar := assert.New(t)
	ar.Equal(22, Backpack9([]int{1, 5, 8, 9, 10, 17, 17, 20}, 8))
	ar.Equal(24, Backpack9([]int{3, 5, 8, 9, 10, 17, 17, 20}, 8))
}
