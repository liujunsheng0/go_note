package algorithm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 回旋镖的数量
// 在平面中给定n个点，每一对点都是不同的，“回旋镖”是一个点的的元组 (i, j, k)，其中 i 和 j 之间的距离与i和k之间的距离相同 （元组的顺序是重要的）。
// 找到回旋镖的数量。 您可以假设n最多为500并且点的坐标都在 [-10000, 10000] （包括）范围内。
// 输入: [[0,0],[1,0],[2,0]]  输出: 2 说明： 两个回旋镖是[[1,0], [0,0], [2,0]]和[[1,0], [2,0], [0,0]]
func NumberOfBoomerangs(points [][]int) int {
	dist := func(p1, p2 []int) int {
		return (p1[0]-p2[0])*(p1[0]-p2[0]) + (p1[1]-p2[1])*(p1[1]-p2[1])
	}
	ans := 0
	// 超时
	//for i := range points {
	//	for j := range points {
	//		for k := range points {
	//			if i == j || i == k || j == k {
	//				continue
	//			}
	//			if dist(points[i], points[j]) == dist(points[i], points[k]) {
	//				ans++
	//			}
	//		}
	//	}
	for i := range points {
		distCount := map[int]int{}
		for j := range points {
			distCount[dist(points[i], points[j])]++
		}
		for _, count := range distCount {
			ans += count * (count - 1)
		}
	}
	return ans
}

func TestNumberOfBoomerangs(t *testing.T) {
	ar := assert.New(t)
	ar.Equal(2, NumberOfBoomerangs([][]int{{0, 0}, {1, 0}, {2, 0}}))
}
